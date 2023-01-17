package interactor

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	datastore "github.com/kod-source/docker-goa-next/app/datastore/mock"
	"github.com/kod-source/docker-goa-next/app/model"
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
