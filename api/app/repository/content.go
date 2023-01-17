package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type ContentRepository interface {
	Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error
}
