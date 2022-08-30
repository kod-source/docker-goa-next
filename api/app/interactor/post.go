package interactor

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/shogo82148/pointer"
)

type PostInteractor interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error)
	ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *string, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error)
	Show(ctx context.Context, id int) (*model.ShowPost, error)
	ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *string, error)
}

type postInteractor struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewPostInteractor(db *sql.DB, tr repository.TimeRepository) PostInteractor {
	return &postInteractor{db: db, tr: tr}
}

func (p *postInteractor) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	var indexPost model.IndexPost
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO posts(`user_id`, `title`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?)",
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := ins.Exec(userID, title, img, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		WHERE p.id = ?
	`, lastID).Scan(
		&indexPost.Post.ID,
		&indexPost.Post.UserID,
		&indexPost.Post.Title,
		&indexPost.Post.Img,
		&indexPost.Post.CreatedAt,
		&indexPost.Post.UpdatedAt,
		&indexPost.User.Name,
		&indexPost.User.Avatar,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &indexPost, nil
}

func (p *postInteractor) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *string, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	rows, err := p.db.Query(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		LEFT JOIN (
			SELECT post_id, COUNT(id) as COUNT
			FROM likes
			GROUP BY post_id
		) as l
		ON p.id = l.post_id
		LEFT JOIN (
			SELECT post_id, COUNT(id) as COUNT
			FROM comments
			GROUP BY post_id
		) as c
		ON p.id = c.post_id
		ORDER BY p.created_at DESC
		LIMIT ?, ?
	`, nextID, limitNumber)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var user model.User
		var countLike *int
		var countComment *int

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Img,
			&post.CreatedAt,
			&post.UpdatedAt,
			&user.Name,
			&user.Avatar,
			&countLike,
			&countComment,
		)
		if err != nil {
			return nil, nil, err
		}

		indexPostsWithCountLike = append(indexPostsWithCountLike, &model.IndexPostWithCountLike{
			IndexPost: model.IndexPost{
				Post: model.Post{
					ID:        post.ID,
					UserID:    post.UserID,
					Title:     post.Title,
					Img:       post.Img,
					CreatedAt: post.CreatedAt,
					UpdatedAt: post.UpdatedAt,
				},
				User: model.User{
					Name:   user.Name,
					Avatar: user.Avatar,
				},
			},
			CountLike:    pointer.IntValue(countLike),
			CountComment: pointer.IntValue(countComment),
		})
	}

	var lastPostID int
	err = p.db.QueryRow(
		"SELECT `id` FROM `posts` ORDER BY `created_at` LIMIT 1",
	).Scan(
		&lastPostID,
	)
	if err != nil {
		return nil, nil, err
	}
	var nextToken *string
	nextToken = pointer.String(fmt.Sprintf("%s/posts?next_id=%d", os.Getenv("END_POINT"), nextID+limitNumber))
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		nextToken = nil
	}

	return indexPostsWithCountLike, nextToken, nil
}

func (p *postInteractor) Delete(ctx context.Context, id int) error {
	stmt, err := p.db.Prepare("DELETE FROM `posts` WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postInteractor) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	upd, err := tx.Prepare("UPDATE `posts` set `title` = ?, `img` = ?, `updated_at` = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = upd.Exec(title, img, p.tr.Now(), id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var indexPost model.IndexPost
	err = tx.QueryRow(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		WHERE p.id = ?
	`, id).Scan(
		&indexPost.Post.ID,
		&indexPost.Post.UserID,
		&indexPost.Post.Title,
		&indexPost.Post.Img,
		&indexPost.Post.CreatedAt,
		&indexPost.Post.UpdatedAt,
		&indexPost.User.Name,
		&indexPost.User.Avatar,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &indexPost, nil
}

func (p *postInteractor) Show(ctx context.Context, id int) (*model.ShowPost, error) {
	var showPost model.ShowPost
	rows, err := p.db.Query(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.id, u.name, u.email, u.created_at, u.avatar, c.id, c.post_id, c.text, c.img, c.created_at, c.updated_at
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		LEFT JOIN comments as c
		ON p.id = c.post_id
		WHERE p.id = ?
		ORDER BY c.created_at DESC
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*model.Comment
	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(
			&showPost.IndexPost.Post.ID,
			&showPost.IndexPost.Post.UserID,
			&showPost.IndexPost.Post.Title,
			&showPost.IndexPost.Post.Img,
			&showPost.IndexPost.Post.CreatedAt,
			&showPost.IndexPost.Post.UpdatedAt,
			&showPost.IndexPost.User.ID,
			&showPost.IndexPost.User.Name,
			&showPost.IndexPost.User.Email,
			&showPost.IndexPost.User.CreatedAt,
			&showPost.IndexPost.User.Avatar,
			&comment.ID,
			&comment.PostID,
			&comment.Text,
			&comment.Img,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	showPost.Comments = comments
	postID := showPost.IndexPost.Post.ID
	if postID == 0 {
		return nil, sql.ErrNoRows
	}

	var likes []*model.Like
	likeRows, err := p.db.Query("SELECT `id`, `user_id`, `post_id` FROM `likes` WHERE `post_id` = ?", postID)
	if err != nil {
		return nil, err
	}
	defer likeRows.Close()
	for likeRows.Next() {
		var like model.Like
		err = likeRows.Scan(
			&like.ID,
			&like.UserID,
			&like.PostID,
		)
		if err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}
	showPost.Likes = likes

	return &showPost, nil
}

func (p *postInteractor) ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *string, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	rows, err := p.db.Query(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT
		FROM posts as p
		INNER JOIN (
			SELECT post_id
			FROM likes
			WHERE user_id = ?
		) as lu
		ON p.id = lu.post_id
		INNER JOIN users as u
		ON p.user_id = u.id
		LEFT JOIN (
			SELECT post_id, COUNT(id) as COUNT
			FROM likes
			GROUP BY post_id
		) as l
		ON p.id = l.post_id
		LEFT JOIN (
			SELECT post_id, COUNT(id) as COUNT
			FROM comments
			GROUP BY post_id
		) as c
		ON p.id = c.post_id
		ORDER BY p.created_at DESC
		LIMIT ?, ?
	`, userID, nextID, limitNumber)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var user model.User
		var countLike *int
		var countComment *int

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Img,
			&post.CreatedAt,
			&post.UpdatedAt,
			&user.Name,
			&user.Avatar,
			&countLike,
			&countComment,
		)
		if err != nil {
			return nil, nil, err
		}

		indexPostsWithCountLike = append(indexPostsWithCountLike, &model.IndexPostWithCountLike{
			IndexPost: model.IndexPost{
				Post: model.Post{
					ID:        post.ID,
					UserID:    post.UserID,
					Title:     post.Title,
					Img:       post.Img,
					CreatedAt: post.CreatedAt,
					UpdatedAt: post.UpdatedAt,
				},
				User: model.User{
					Name:   user.Name,
					Avatar: user.Avatar,
				},
			},
			CountLike:    pointer.IntValue(countLike),
			CountComment: pointer.IntValue(countComment),
		})
	}

	var lastPostID int
	err = p.db.QueryRow(`
		SELECT p.id
		FROM posts AS p
		INNER JOIN (
			SELECT post_id
			FROM likes
			WHERE user_id = ?
		) AS l
		ON p.id = l.post_id
		ORDER BY p.created_at
		LIMIT 1
	`, userID).Scan(
		&lastPostID,
	)
	var nextToken *string
	nextToken = pointer.String(fmt.Sprintf("%s/posts?next_id=%d", os.Getenv("END_POINT"), nextID+limitNumber))
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		nextToken = nil
	}

	return indexPostsWithCountLike, nextToken, nil
}
