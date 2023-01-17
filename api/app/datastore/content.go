package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.ContentRepository = (*contentDatastore)(nil)

var ContentRepositorySet = wire.NewSet(
	NewContentRepository,
	wire.Bind(new(repository.ContentRepository), new(*contentDatastore)),
)

type contentDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewContentRepository(db *sql.DB, tr repository.TimeRepository) *contentDatastore {
	return &contentDatastore{db: db, tr: tr}
}

func (cd *contentDatastore) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	return nil
}
