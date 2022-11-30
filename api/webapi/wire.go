//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/datastore"
	"github.com/kod-source/docker-goa-next/app/external"
	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
)

func NewApp(ctx context.Context) (*App, error) {
	wire.Build(
		// Application
		newApp, newService,
		// DB
		schema.NewDB,
		// TimeRepostiory
		repository.TimeRepositorySet,
		// service (external)
		external.Set,
		// repostiory (datastore) ...
		datastore.Set,
		// usecase (interactor) ...
		interactor.Set,
		// Controller ...
		ControllerSet,
	)

	return &App{}, nil
}
