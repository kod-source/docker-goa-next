package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
)

func Test_CreatePost(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
	}
	pd := NewPostDatastore(testDB, tr)

	t.Run("[OK]投稿作成", func(t *testing.T) {
		got, err := pd.CreatePost(ctx, 1, "create_post", pointer.Ptr("create_img_post"))
		if err != nil {
			t.Fatal(err)
		}

		post, err := schema.SelectPost(ctx, testDB, &schema.Post{ID: uint64(got.Post.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.IndexPost{
			Post: model.Post{
				ID:        int(post.ID),
				UserID:    int(post.UserID),
				Title:     post.Title,
				Img:       pointer.PtrOrNil(post.Img.String),
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			},
			User: model.User{
				ID:     model.UserID(user.ID),
				Name:   user.Name,
				Avatar: pointer.PtrOrNil(user.Avatar.String),
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := pd.Delete(ctx, got.Post.ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]投稿作成 - 画像がnilの時", func(t *testing.T) {
		got, err := pd.CreatePost(ctx, 1, "create_post", nil)
		if err != nil {
			t.Fatal(err)
		}

		post, err := schema.SelectPost(ctx, testDB, &schema.Post{ID: uint64(got.Post.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.IndexPost{
			Post: model.Post{
				ID:        int(post.ID),
				UserID:    int(post.UserID),
				Title:     post.Title,
				Img:       nil,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			},
			User: model.User{
				ID:     model.UserID(user.ID),
				Name:   user.Name,
				Avatar: pointer.PtrOrNil(user.Avatar.String),
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := pd.Delete(ctx, got.Post.ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]投稿作成 - 存在しないUserIDを指定した時", func(t *testing.T) {
		_, err := pd.CreatePost(ctx, 1000, "test", pointer.Ptr("test"))
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error %v, but got error is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}

func Test_ShowAll(t *testing.T) {
	pd := NewPostDatastore(testDB, nil)

	t.Run("[OK]全ての投稿データを取得", func(t *testing.T) {
		want := []*model.IndexPostWithCountLike{
			{
				IndexPost: model.IndexPost{
					Post: model.Post{
						ID:        2,
						UserID:    1,
						Title:     "test2_title",
						Img:       pointer.Ptr("test2_post_img"),
						CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
					},
					User: model.User{
						Name:   "test1_name",
						Avatar: pointer.Ptr("test1_avatar"),
					},
				},
				CountLike:    1,
				CountComment: 1,
			},
			{
				IndexPost: model.IndexPost{
					Post: model.Post{
						ID:        3,
						UserID:    2,
						Title:     "test3_title",
						Img:       nil,
						CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					},
					User: model.User{
						Name:   "test2_name",
						Avatar: nil,
					},
				},
				CountLike:    0,
				CountComment: 0,
			},
			{
				IndexPost: model.IndexPost{
					Post: model.Post{
						ID:        1,
						UserID:    1,
						Title:     "test1_title",
						Img:       pointer.Ptr("test1_post_img"),
						CreatedAt: now,
						UpdatedAt: now,
					},
					User: model.User{
						Name:   "test1_name",
						Avatar: pointer.Ptr("test1_avatar"),
					},
				},
				CountLike:    1,
				CountComment: 2,
			},
		}

		got, nextID, err := pd.ShowAll(ctx, 0)
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("want nextID is nil but got is %d", *nextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]全ての投稿データを取得 - nextIDを指定", func(t *testing.T) {
		want := []*model.IndexPostWithCountLike{
			{
				IndexPost: model.IndexPost{
					Post: model.Post{
						ID:        1,
						UserID:    1,
						Title:     "test1_title",
						Img:       pointer.Ptr("test1_post_img"),
						CreatedAt: now,
						UpdatedAt: now,
					},
					User: model.User{
						Name:   "test1_name",
						Avatar: pointer.Ptr("test1_avatar"),
					},
				},
				CountLike:    1,
				CountComment: 2,
			},
		}

		got, nextID, err := pd.ShowAll(ctx, 2)
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("want nextID is nil but got is %d", *nextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]全ての投稿データを取得 - データがない時", func(t *testing.T) {
		var want []*model.IndexPostWithCountLike

		got, nextID, err := pd.ShowAll(ctx, 1000)
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("want nextID is nil but got is %d", *nextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]全ての投稿データを取得 - 全てのデータをとりきれない", func(t *testing.T) {
		var posts []*schema.Post
		for i := 0; i < 20; i++ {
			id := i + 7
			posts = append(posts, &schema.Post{
				ID:        uint64(id),
				UserID:    2,
				Title:     fmt.Sprintf("create_title_%d", id),
				CreatedAt: time.Date(2022, 3, id, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 3, id, 0, 0, 0, 0, jst),
				Img: sql.NullString{
					String: fmt.Sprintf("create_img_%d", id),
					Valid:  false,
				},
			})
		}
		if err := schema.InsertPost(ctx, testDB, posts...); err != nil {
			t.Fatal(err)
		}

		wantNextID := 7
		_, nextID, err := pd.ShowAll(ctx, 0)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(pointer.Ptr(wantNextID), nextID); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}

		for _, p := range posts {
			if err := pd.Delete(ctx, int(p.ID)); err != nil {
				t.Fatal(err)
			}
		}
	})
}

func Test_DeletePost(t *testing.T) {
	pd := NewPostDatastore(testDB, nil)

	t.Run("[OK]投稿削除", func(t *testing.T) {
		postID := 27
		p := &schema.Post{
			ID:        uint64(postID),
			UserID:    2,
			Title:     "delete_post",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "delete_img",
				Valid:  true,
			},
		}
		if err := schema.InsertPost(ctx, testDB, p); err != nil {
			t.Fatal(err)
		}
		got, err := schema.SelectPost(ctx, testDB, &schema.Post{ID: uint64(postID)})
		if err != nil {
			t.Fatal(err)
		}
		commentID := 9
		c := &schema.Comment{
			ID:        9,
			PostID:    uint64(postID),
			UserID:    1,
			Text:      "delete_ctext",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "delete_img",
				Valid:  true,
			},
		}
		if err := schema.InsertComment(ctx, testDB, c); err != nil {
			t.Fatal(err)
		}
		gotComment, err := schema.SelectComment(ctx, testDB, &schema.Comment{ID: uint64(commentID)})
		if err != nil {
			t.Fatal(err)
		}

		if err := pd.Delete(ctx, postID); err != nil {
			t.Fatal(err)
		}

		if _, err := schema.SelectPost(ctx, testDB, &schema.Post{ID: got.ID}); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v but got error is %v", sql.ErrNoRows, err)
		}
		if _, err := schema.SelectComment(ctx, testDB, &schema.Comment{ID: gotComment.ID}); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v but got error is %v", sql.ErrNoRows, err)
		}
	})
}

func Test_UpdatePost(t *testing.T) {
	updateNow := time.Date(2022, 1, 31, 0, 0, 0, 0, jst)
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return updateNow
	}
	pd := NewPostDatastore(testDB, tr)

	t.Run("[OK]投稿の更新", func(t *testing.T) {
		postID := 28
		p := &schema.Post{
			ID:        uint64(postID),
			UserID:    2,
			Title:     "create_title",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "create_img",
				Valid:  true,
			},
		}
		if err := schema.InsertPost(ctx, testDB, p); err != nil {
			t.Fatal(err)
		}

		want := &model.IndexPost{
			Post: model.Post{
				ID:        postID,
				UserID:    2,
				Title:     "updated_title",
				Img:       pointer.Ptr("updated_img"),
				CreatedAt: now,
				UpdatedAt: updateNow,
			},
			User: model.User{
				Name:   "test2_name",
				Avatar: nil,
			},
		}

		got, err := pd.Update(ctx, postID, "updated_title", pointer.Ptr("updated_img"))
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := pd.Delete(ctx, postID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]投稿の更新 - 画像をnilに更新", func(t *testing.T) {
		postID := 29
		p := &schema.Post{
			ID:        uint64(postID),
			UserID:    1,
			Title:     "create_title",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "create_img",
				Valid:  true,
			},
		}
		if err := schema.InsertPost(ctx, testDB, p); err != nil {
			t.Fatal(err)
		}

		want := &model.IndexPost{
			Post: model.Post{
				ID:        postID,
				UserID:    1,
				Title:     "updated_title",
				Img:       nil,
				CreatedAt: now,
				UpdatedAt: updateNow,
			},
			User: model.User{
				Name:   "test1_name",
				Avatar: pointer.Ptr("test1_avatar"),
			},
		}

		got, err := pd.Update(ctx, postID, "updated_title", nil)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := pd.Delete(ctx, postID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]投稿の更新 - 存在しないIDを指定した時", func(t *testing.T) {
		if _, err := pd.Update(ctx, 1000, "update_title", nil); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}
