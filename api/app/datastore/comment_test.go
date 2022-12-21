package datastore

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
)

func Test_Create(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
	}
	cd := NewCommentDatastore(testDB, tr)

	t.Run("[OK]コメント作成", func(t *testing.T) {
		got, err := cd.Create(ctx, 2, 1, "create_comment", pointer.Ptr("create_img"))
		if err != nil {
			t.Fatal(err)
		}

		comment, err := schema.SelectComment(ctx, testDB, &schema.Comment{ID: uint64(got.Comment.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.CommentWithUser{
			Comment: model.Comment{
				ID:        int(comment.ID),
				PostID:    int(comment.PostID),
				UserID:    int(comment.UserID),
				Text:      comment.Text,
				Img:       pointer.PtrOrNil(comment.Img.String),
				CreatedAt: pointer.PtrOrNil(comment.CreatedAt),
				UpdatedAt: pointer.PtrOrNil(comment.UpdatedAt),
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
		if err := cd.Delete(ctx, got.Comment.ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]コメント作成 - 画像がnilの時", func(t *testing.T) {
		got, err := cd.Create(ctx, 2, 1, "create_comment", nil)
		if err != nil {
			t.Fatal(err)
		}

		comment, err := schema.SelectComment(ctx, testDB, &schema.Comment{ID: uint64(got.Comment.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.CommentWithUser{
			Comment: model.Comment{
				ID:        int(comment.ID),
				PostID:    int(comment.PostID),
				UserID:    int(comment.UserID),
				Text:      comment.Text,
				Img:       pointer.PtrOrNil(comment.Img.String),
				CreatedAt: pointer.PtrOrNil(comment.CreatedAt),
				UpdatedAt: pointer.PtrOrNil(comment.UpdatedAt),
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
		if err := cd.Delete(ctx, got.Comment.ID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]コメント作成 - 存在しないPostIDを指定した時", func(t *testing.T) {
		_, err := cd.Create(ctx, 1000, 1, "create_comment", pointer.Ptr("create_img"))
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error code %v, but got err code is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})

	t.Run("[NG]コメント作成 - 存在しないUserIDを指定した時", func(t *testing.T) {
		_, err := cd.Create(ctx, 2, 1000, "create_comment", pointer.Ptr("create_img"))
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error code %v, but got err code is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}

func Test_ShowByPostID(t *testing.T) {
	cd := NewCommentDatastore(testDB, nil)

	t.Run("[OK]投稿に紐づくコメント取得", func(t *testing.T) {
		want := []*model.CommentWithUser{
			{
				Comment: model.Comment{
					ID:        2,
					PostID:    1,
					UserID:    2,
					Text:      "test2_comment",
					Img:       nil,
					CreatedAt: pointer.Ptr(time.Date(2022, 2, 1, 0, 0, 0, 0, jst)),
					UpdatedAt: pointer.Ptr(time.Date(2022, 2, 1, 0, 0, 0, 0, jst)),
				},
				User: model.User{
					ID:     2,
					Name:   "test2_name",
					Avatar: nil,
				},
			},
			{
				Comment: model.Comment{
					ID:        1,
					PostID:    1,
					UserID:    1,
					Text:      "test1_comment",
					Img:       pointer.PtrOrNil("test1_comment_img"),
					CreatedAt: pointer.Ptr(now),
					UpdatedAt: pointer.Ptr(now),
				},
				User: model.User{
					ID:     1,
					Name:   "test1_name",
					Avatar: pointer.PtrOrNil("test1_avatar"),
				},
			},
		}

		got, err := cd.ShowByPostID(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]投稿に紐づくコメント取得 - 投稿は存在するがコメントが存在しない時", func(t *testing.T) {
		if _, err := cd.ShowByPostID(ctx, 3); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
	t.Run("[NG]投稿に紐づくコメント取得 - 存在しない投稿を指定した時", func(t *testing.T) {
		if _, err := cd.ShowByPostID(ctx, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}

func Test_Update(t *testing.T) {
	updatedTime := time.Date(2022, 2, 1, 0, 0, 0, 0, jst)
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return updatedTime
	}
	cd := NewCommentDatastore(testDB, tr)

	t.Run("[OK]コメントの更新", func(t *testing.T) {
		want := &model.Comment{
			ID:        3,
			PostID:    2,
			UserID:    1,
			Text:      "update_text",
			Img:       pointer.Ptr("update_img"),
			CreatedAt: &now,
			UpdatedAt: &updatedTime,
		}

		got, err := cd.Update(ctx, 3, "update_text", pointer.Ptr("update_img"))
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]コメントの更新 - 画像をnilにする", func(t *testing.T) {
		want := &model.Comment{
			ID:        3,
			PostID:    2,
			UserID:    1,
			Text:      "update_text",
			Img:       nil,
			CreatedAt: &now,
			UpdatedAt: &updatedTime,
		}

		got, err := cd.Update(ctx, 3, "update_text", nil)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]コメント更新 - 存在しないIDを指定した時", func(t *testing.T) {
		if _, err := cd.Update(ctx, 1000, "update_test", pointer.Ptr("update_img")); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}

func Test_Delete(t *testing.T) {
	cd := NewCommentDatastore(testDB, nil)

	t.Run("[OK]コメント削除", func(t *testing.T) {
		commentID := 8
		comment := &schema.Comment{
			ID:        uint64(commentID),
			PostID:    2,
			UserID:    2,
			Text:      "delete_text",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "delete_img",
				Valid:  true,
			},
		}
		if err := schema.InsertComment(ctx, testDB, comment); err != nil {
			t.Fatal(err)
		}

		if err := cd.Delete(ctx, commentID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]コメント削除 - 存在しないIDを指定した時", func(t *testing.T) {
		if err := cd.Delete(ctx, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}
