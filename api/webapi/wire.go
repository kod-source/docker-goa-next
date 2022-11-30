//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/datastore"
	"github.com/kod-source/docker-goa-next/app/external"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

func NewApp(ctx context.Context) (*App, error) {
	wire.Build(
		// Application
		newApp, newService,
		// DB
		schema.NewDB,
		// TimeRepostiory
		repository.TimeRepositorySet,
		// external
		external.Set,
		// datastore ...
		datastore.Set,
		// usecase ...
		usecase.Set,
		// Controller ...
		ControllerSet,
	)

	return &App{}, nil
}
