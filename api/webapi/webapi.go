package main

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

type App struct {
	srv *goa.Service
}

func newService() *goa.Service {
	// Create service
	srv := goa.New("docker_goa_next")

	// Mount middleware
	srv.Use(middleware.RequestID())
	srv.Use(middleware.LogRequest(true))
	srv.Use(middleware.ErrorHandler(srv, true))
	srv.Use(middleware.Recover())

	// 認証付きAPIの設定
	app.UseJWTMiddleware(srv, newAuthMiddleware())
	return srv
}

func newApp(ctx context.Context, srv *goa.Service, userCtrl app.UsersController, postCtrl app.PostsController, opeCtrl app.OperandsController, likeCtrl app.LikesController, commentCtrl app.CommentsController, authCtrl app.AuthController) (*App, error) {
	app.MountOperandsController(srv, opeCtrl)
	app.MountAuthController(srv, authCtrl)
	app.MountUsersController(srv, userCtrl)
	app.MountPostsController(srv, postCtrl)
	app.MountCommentsController(srv, commentCtrl)
	app.MountLikesController(srv, likeCtrl)

	return &App{srv: srv}, nil
}

// ControllerSet ...
var ControllerSet = wire.NewSet(
	NewUsersController, wire.Bind(new(app.UsersController), new(*UsersController)),
	NewPostsController, wire.Bind(new(app.PostsController), new(*PostsController)),
	NewOperandsController, wire.Bind(new(app.OperandsController), new(*OperandsController)),
	NewLikesController, wire.Bind(new(app.LikesController), new(*LikesController)),
	NewCommentsController, wire.Bind(new(app.CommentsController), new(*CommentsController)),
	NewAuthController, wire.Bind(new(app.AuthController), new(*AuthController)),
)
