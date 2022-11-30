package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/shogo82148/pointer"
)

var _ PostDatastore = (*postDatastore)(nil)

var PostDatastoreSet = wire.NewSet(
	NewPostDatastore,
	wire.Bind(new(PostDatastore), new(*postDatastore)),
)

type PostDatastore interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error)
	ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error)
	Show(ctx context.Context, id int) (*model.ShowPost, error)
	ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
	// ShowPostMy 指定したUserIDが投稿したものを取得する
	ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
	// ShowPostMedia 指定したUserIDの画像投稿したものを取得する
	ShowPostMedia(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error)
}

type postDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewPostDatastore(db *sql.DB, tr repository.TimeRepository) *postDatastore {
	return &postDatastore{db: db, tr: tr}
}

func (p *postDatastore) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	var indexPost model.IndexPost
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO post(`user_id`, `title`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?)",
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := ins.Exec(userID, title, img, p.tr.Now(), p.tr.Now())
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
		FROM post as p
		INNER JOIN user as u
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

func (p *postDatastore) ShowAll(ctx context.Context, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	rows, err := p.db.Query("SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT"+
		" FROM `post` AS `p`"+
		" INNER JOIN `user` AS `u`"+
		" ON p.user_id = u.id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `like`"+
		" GROUP BY `post_id`"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `comment`"+
		" GROUP BY `post_id`"+
		" ) AS `c`"+
		" ON p.id = c.post_id"+
		" ORDER BY p.created_at DESC"+
		" LIMIT ?, ?", nextID, limitNumber)
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
		"SELECT `id` FROM `post` ORDER BY `created_at` LIMIT 1",
	).Scan(
		&lastPostID,
	)
	if err != nil {
		return nil, nil, err
	}
	var resNextID *int
	resNextID = pointer.Int(nextID + limitNumber)
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		resNextID = nil
	}

	return indexPostsWithCountLike, resNextID, nil
}

func (p *postDatastore) Delete(ctx context.Context, id int) error {
	stmt, err := p.db.Prepare("DELETE FROM `post` WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (p *postDatastore) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	upd, err := tx.Prepare("UPDATE `post` set `title` = ?, `img` = ?, `updated_at` = ? WHERE id = ?")
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
		FROM post as p
		INNER JOIN user as u
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

func (p *postDatastore) Show(ctx context.Context, id int) (*model.ShowPost, error) {
	var showPost model.ShowPost
	rows, err := p.db.Query("SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.id, u.name, u.email, u.created_at, u.avatar, cu.C_ID, cu.C_POST_ID, cu.C_TEXT, cu.C_IMG, cu.C_CREATED_AT, cu.C_UPDATED_AT, cu.U_ID, cu.U_NAME, cu.U_AVATAR"+
		" FROM `post` as `p`"+
		" INNER JOIN `user` as `u`"+
		" ON p.user_id = u.id"+
		" LEFT JOIN"+
		"("+
		" SELECT c.id as C_ID, c.post_id as C_POST_ID, c.text as C_TEXT, c.img as C_IMG, c.created_at as C_CREATED_AT, c.updated_at as C_UPDATED_AT, u.id as U_ID, u.name as U_NAME, u.avatar as U_AVATAR"+
		" FROM `comment` as `c`"+
		" INNER JOIN `user` as `u`"+
		" ON c.user_id = u.id"+
		" ) as `cu`"+
		" ON p.id = cu.C_POST_ID"+
		" WHERE p.id = ?"+
		" ORDER BY cu.C_CREATED_AT DESC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentsWithUsers []*model.ShowCommentWithUser
	for rows.Next() {
		var commentWithUser model.ShowCommentWithUser

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
			&commentWithUser.Comment.ID,
			&commentWithUser.Comment.PostID,
			&commentWithUser.Comment.Text,
			&commentWithUser.Comment.Img,
			&commentWithUser.Comment.CreatedAt,
			&commentWithUser.Comment.UpdatedAt,
			&commentWithUser.User.ID,
			&commentWithUser.User.Name,
			&commentWithUser.User.Avatar,
		)
		if err != nil {
			return nil, err
		}
		commentsWithUsers = append(commentsWithUsers, &commentWithUser)
	}
	showPost.CommenstWithUsers = commentsWithUsers
	postID := showPost.IndexPost.Post.ID
	if postID == 0 {
		return nil, sql.ErrNoRows
	}

	var likes []*model.Like
	likeRows, err := p.db.Query("SELECT `id`, `user_id`, `post_id` FROM `like` WHERE `post_id` = ?", postID)
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

func (p *postDatastore) ShowMyLike(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	rows, err := p.db.Query("SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT"+
		" FROM `post` AS `p`"+
		" INNER JOIN ("+
		" SELECT `post_id`"+
		" FROM `like`"+
		" WHERE `user_id` = ?"+
		" ) AS `lu`"+
		" ON p.id = lu.post_id"+
		" INNER JOIN `user` AS `u`"+
		" ON p.user_id = u.id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `like`"+
		" GROUP BY `post_id`"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `comment`"+
		" GROUP BY `post_id`"+
		" ) AS `c`"+
		" ON p.id = c.post_id"+
		" ORDER BY p.created_at DESC"+
		" LIMIT ?, ?", userID, nextID, limitNumber)
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
	err = p.db.QueryRow("SELECT p.id"+
		" FROM `post` AS `p`"+
		" INNER JOIN"+
		" ("+
		" SELECT `post_id`"+
		" FROM `like`"+
		" WHERE `user_id` = ?"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" ORDER BY p.created_at"+
		" LIMIT 1", userID).Scan(
		&lastPostID,
	)
	var resNextID *int
	resNextID = pointer.Int(nextID + limitNumber)
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		resNextID = nil
	}

	return indexPostsWithCountLike, resNextID, nil
}

// ShowPostMy 指定したUserIDが投稿したものを取得する
func (p *postDatastore) ShowPostMy(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	tx, err := p.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT"+
		" FROM `post` AS `p`"+
		" INNER JOIN `user` AS `u`"+
		" ON p.user_id = u.id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `like`"+
		" GROUP BY `post_id`"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `comment`"+
		" GROUP BY `post_id`"+
		" ) AS `c`"+
		" ON p.id = c.post_id"+
		" WHERE p.user_id = ?"+
		" ORDER BY p.created_at DESC"+
		" LIMIT ?, ?", userID, nextID, limitNumber)
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
	err = tx.QueryRow(`
		SELECT p.id
		FROM post AS p
		INNER JOIN (
			SELECT post_id
			FROM like
			WHERE user_id = ?
		) AS l
		ON p.id = l.post_id
		ORDER BY p.created_at
		LIMIT 1
	`, userID).Scan(
		&lastPostID,
	)
	var resNextID *int
	resNextID = pointer.Int(nextID + limitNumber)
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		resNextID = nil
	}
	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return indexPostsWithCountLike, resNextID, nil
}

func (p *postDatastore) ShowPostMedia(ctx context.Context, userID, nextID int) ([]*model.IndexPostWithCountLike, *int, error) {
	var indexPostsWithCountLike []*model.IndexPostWithCountLike
	limitNumber := 20
	tx, err := p.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar, l.COUNT, c.COUNT"+
		" FROM `post` AS `p`"+
		" INNER JOIN `user` AS `u`"+
		" ON p.user_id = u.id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `like`"+
		" GROUP BY `post_id`"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" LEFT JOIN"+
		" ("+
		" SELECT `post_id`, COUNT(id) AS `COUNT`"+
		" FROM `comment`"+
		" GROUP BY `post_id`"+
		" ) AS `c`"+
		" ON p.id = c.post_id"+
		" WHERE p.user_id = ? AND p.img IS NOT NULL AND p.img != ''"+
		" ORDER BY p.created_at DESC"+
		" LIMIT ?, ?", userID, nextID, limitNumber)
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
	err = tx.QueryRow("SELECT p.id"+
		" FROM `post` AS `p`"+
		" INNER JOIN ("+
		" SELECT `post_id`"+
		" FROM `like`"+
		" WHERE `user_id` = ?"+
		" ) AS `l`"+
		" ON p.id = l.post_id"+
		" ORDER BY p.created_at"+
		" LIMIT 1", userID).Scan(
		&lastPostID,
	)
	var resNextID *int
	resNextID = pointer.Int(nextID + limitNumber)
	if len(indexPostsWithCountLike) == 0 || indexPostsWithCountLike[len(indexPostsWithCountLike)-1].IndexPost.Post.ID == lastPostID {
		resNextID = nil
	}
	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}

	return indexPostsWithCountLike, resNextID, nil
}
