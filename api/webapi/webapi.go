package main

import (
	"context"
	"net"

	"github.com/caarlos0/env"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

type App struct {
	srv *goa.Service
}

type appConfig struct {
	ApiEndPoint      string `env:"END_POINT"`
	DatabaseName     string `env:"MYSQL_DATABASE,required"`
	DatabaseUser     string `env:"MYSQL_USER,required"`
	DatabasePassword string `env:"MYSQL_PASSWORD,required"`
	DatabasePort     string `env:"MYSQL_PORT" envDefault:"3306"`
	DatabaseHost     string `env:"MYSQL_HOST,required"`
}

func getAppConfig() (*appConfig, error) {
	cfg := appConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
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

func newMysqlConfig(conf *appConfig) (*mysql.Config, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = conf.DatabaseUser
	config.Passwd = conf.DatabasePassword
	config.Addr = net.JoinHostPort(conf.DatabaseHost, conf.DatabasePort)
	config.DBName = conf.DatabaseName
	config.ParseTime = true

	return config, nil
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
