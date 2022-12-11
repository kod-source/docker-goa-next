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
)

func NewApp(ctx context.Context) (*App, error) {
	wire.Build(
		// Application
		getAppConfig, newApp, newService, newMysqlConfig,
		// DB
		datastore.NewDB,
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
