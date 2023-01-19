package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	interactor "github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
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
