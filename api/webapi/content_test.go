package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	interactor "github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
	"github.com/shogo82148/pointer"
)

func Test_DeleteContent(t *testing.T) {
	srv := testApp.srv
	cu := &interactor.MockContentUsecase{}
	cc := NewContentController(srv, cu)

	wantContentID := model.ContentID(1)
	wantUserID := model.UserID(2)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantUserID))

	t.Run("[OK]コンテントの削除", func(t *testing.T) {
		cu.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantContentID, contentID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil
		}
		defer func() {
			cu.DeleteFunc = nil
		}()

		test.DeleteContentOK(t, ctx, srv, cc, int(wantContentID))
	})

	t.Run("[NG]コンテントの削除 - 権限エラーの時", func(t *testing.T) {
		cu.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantContentID, contentID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return myerrors.ErrBadRequestNoPermission
		}
		defer func() {
			cu.DeleteFunc = nil
		}()

		test.DeleteContentBadRequest(t, ctx, srv, cc, int(wantContentID))
	})

	t.Run("[NG]コンテントの削除 - 不明なIDを指定した時", func(t *testing.T) {
		cu.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return sql.ErrNoRows
		}
		defer func() {
			cu.DeleteFunc = nil
		}()

		test.DeleteContentNotFound(t, ctx, srv, cc, 1000)
	})

	t.Run("[NG]コンテントの削除 - 想定外エラー発生", func(t *testing.T) {
		cu.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantContentID, contentID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return errors.New("test error")
		}
		defer func() {
			cu.DeleteFunc = nil
		}()

		test.DeleteContentInternalServerError(t, ctx, srv, cc, int(wantContentID))
	})
}

func Test_CreateContent(t *testing.T) {
	srv := testApp.srv
	cu := &interactor.MockContentUsecase{}
	cc := NewContentController(srv, cu)

	wantContentID := model.ContentID(1)
	wantThreadID := model.ThreadID(3)
	wantText := "create content"
	wantImg := "content img"
	wantUserID := model.UserID(2)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantUserID))

	t.Run("[OK]コンテント作成", func(t *testing.T) {
		mcu := &model.ContentUser{
			Content: model.Content{
				ID:        wantContentID,
				UserID:    wantUserID,
				ThreadID:  wantThreadID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       pointer.Ptr(wantImg),
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "user1",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("avatar"),
			},
		}
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, *img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return mcu, nil
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		want := &app.ContentUser{
			Content: &app.Content{
				ID:        int(wantContentID),
				UserID:    int(wantUserID),
				ThreadID:  int(wantThreadID),
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       pointer.Ptr(wantImg),
			},
			User: &app.ShowUser{
				ID:        int(wantUserID),
				Name:      "user1",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("avatar"),
			},
		}

		_, got := test.CreateContentCreated(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      pointer.Ptr(wantImg),
			Text:     wantText,
			ThreadID: int(wantThreadID),
		})
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)`\n%s", diff)
		}
	})

	t.Run("[OK]コンテント作成 - 画像がnilの時", func(t *testing.T) {
		mcu := &model.ContentUser{
			Content: model.Content{
				ID:        wantContentID,
				UserID:    wantUserID,
				ThreadID:  wantThreadID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "user1",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    nil,
			},
		}
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if img != nil {
				t.Errorf("img mismatch (-want nil, +got %s)", *img)
			}

			return mcu, nil
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		want := &app.ContentUser{
			Content: &app.Content{
				ID:        int(wantContentID),
				UserID:    int(wantUserID),
				ThreadID:  int(wantThreadID),
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
			},
			User: &app.ShowUser{
				ID:        int(wantUserID),
				Name:      "user1",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    nil,
			},
		}

		_, got := test.CreateContentCreated(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      nil,
			Text:     wantText,
			ThreadID: int(wantThreadID),
		})
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)`\n%s", diff)
		}
	})

	t.Run("[NG]コンテント作成 - 外部キーエラーの時", func(t *testing.T) {
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, *img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.MySQLErrorAddOrUpdateForeignKey
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		test.CreateContentBadRequest(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      &wantImg,
			Text:     wantText,
			ThreadID: int(wantThreadID),
		})
	})

	t.Run("[NG]コンテント作成 - 不明な数字の時", func(t *testing.T) {
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, *img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.ErrBadRequestInt
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		test.CreateContentBadRequest(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      &wantImg,
			Text:     wantText,
			ThreadID: int(0),
		})
	})

	t.Run("[NG]コンテント作成 - 空文字の時", func(t *testing.T) {
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, *img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.ErrBadRequestSting
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		test.CreateContentBadRequest(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      &wantImg,
			Text:     "",
			ThreadID: int(wantThreadID),
		})
	})

	t.Run("[NG]コンテント作成 - 想定外エラー発生", func(t *testing.T) {
		cu.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, *img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, errors.New("test error")
		}
		defer func() {
			cu.CreateFunc = nil
		}()

		test.CreateContentInternalServerError(t, ctx, srv, cc, &app.CreateContentPayload{
			Img:      &wantImg,
			Text:     wantText,
			ThreadID: int(wantThreadID),
		})
	})
}
