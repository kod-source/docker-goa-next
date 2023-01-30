package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	datastore "github.com/kod-source/docker-goa-next/app/datastore/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/shogo82148/pointer"
)

func Test_DeleteContent(t *testing.T) {
	cr := &datastore.MockContentRepository{}
	cu := NewContentUsecase(cr)
	wantContentID := model.ContentID(1)
	wantUserID := model.UserID(2)

	t.Run("[OK]コンテントの削除", func(t *testing.T) {
		cr.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantContentID, contentID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil
		}

		if err := cu.Delete(ctx, wantUserID, wantContentID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]コンテントの削除 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		cr.DeleteFunc = func(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantContentID, contentID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return wantErr
		}

		if err := cu.Delete(ctx, wantUserID, wantContentID); !errors.Is(err, wantErr) {
			t.Errorf("error mismatch (-want %v, +got %v)", wantErr, err)
		}
	})
}

func Test_CreateContent(t *testing.T) {
	cr := &datastore.MockContentRepository{}
	cu := NewContentUsecase(cr)
	wantContentID := model.ContentID(1)
	wantUserID := model.UserID(2)
	wantThreadID := model.ThreadID(3)
	wantText := "create content"
	wantImg := "content img"

	t.Run("[OK]コンテントの作成", func(t *testing.T) {
		want := &model.ContentUser{
			Content: model.Content{
				ID:        wantContentID,
				UserID:    wantUserID,
				ThreadID:  wantThreadID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       &wantImg,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "user1",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("avatar"),
			},
		}
		cr.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return want, nil
		}
		defer func() {
			cr.CreateFunc = nil
		}()

		got, err := cu.Create(ctx, wantText, wantThreadID, wantUserID, &wantImg)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]コンテントの作成 - 空文字の時", func(t *testing.T) {
		if _, err := cu.Create(ctx, "", wantThreadID, wantUserID, &wantImg); !errors.Is(err, myerrors.ErrBadRequestSting) {
			t.Errorf("error mismatch (-want %v, +got %v)", myerrors.ErrBadRequestSting, err)
		}
	})

	t.Run("[NG]コンテントの作成 - スレッドIDが0の時", func(t *testing.T) {
		if _, err := cu.Create(ctx, wantText, 0, wantUserID, &wantImg); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error mismatch (-want %v, +got %v)", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]コンテントの作成 - ユーザーIDが0の時", func(t *testing.T) {
		if _, err := cu.Create(ctx, wantText, wantThreadID, model.UserID(0), &wantImg); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error mismatch (-want %v, +got %v)", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]コンテントの作成 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		cr.CreateFunc = func(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, wantErr
		}
		defer func() {
			cr.CreateFunc = nil
		}()

		if _, err := cu.Create(ctx, wantText, wantThreadID, wantUserID, &wantImg); !errors.Is(err, wantErr) {
			t.Errorf("error mismatch (-want %v, +got %v)", wantErr, err)
		}
	})
}
