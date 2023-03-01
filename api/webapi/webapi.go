//go:generate go run gen/main.go

package webapi

import (
	"context"
	"fmt"
	"net"

	"github.com/caarlos0/env"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

type App struct {
	srv *goa.Service
}

type appConfig struct {
	APIEndPoint      string `env:"END_POINT"`
	ClientEndPoint   string `env:"CLIENT_END_POINT"`
	DatabaseName     string `env:"MYSQL_DATABASE,required"`
	DatabaseUser     string `env:"MYSQL_USER,required"`
	DatabasePassword string `env:"MYSQL_PASSWORD,required"`
	DatabasePort     string `env:"MYSQL_PORT" envDefault:"3306"`
	DatabaseHost     string `env:"MYSQL_HOST,required"`
}

type GoogleConfig struct {
	ClientEndPoint string `env:"CLIENT_END_POINT"`
	ClientID       string `env:"GOOGLE_CLIENT_ID"`
	ClientSecret   string `env:"GOOGLE_CLIENT_SECRET"`
}

func (a *App) LaunchServer() error {
	if err := a.srv.ListenAndServe(":3000"); err != nil {
		a.srv.LogError("startup", "err", err)
		return err
	}
	return nil
}

func getAppConfig() (*appConfig, error) {
	cfg := appConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func newGoogleConfig() (*oauth2.Config, error) {
	gcfg := GoogleConfig{}
	if err := env.Parse(&gcfg); err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     gcfg.ClientID,
		ClientSecret: gcfg.ClientSecret,
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       []string{"openid"},
		RedirectURL:  fmt.Sprintf("%s/auth/callback/google", gcfg.ClientEndPoint),
	}, nil
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

func newApp(ctx context.Context, srv *goa.Service, userCtrl app.UsersController, postCtrl app.PostsController, opeCtrl app.OperandsController, likeCtrl app.LikesController, commentCtrl app.CommentsController, authCtrl app.AuthController, roomCtrl app.RoomsController, userRoomCtrl app.UserRoomsController, threadCtrl app.ThreadsController, contentCtrl app.ContentController) (*App, error) {
	app.MountOperandsController(srv, opeCtrl)
	app.MountAuthController(srv, authCtrl)
	app.MountUsersController(srv, userCtrl)
	app.MountPostsController(srv, postCtrl)
	app.MountCommentsController(srv, commentCtrl)
	app.MountLikesController(srv, likeCtrl)
	app.MountRoomsController(srv, roomCtrl)
	app.MountUserRoomsController(srv, userRoomCtrl)
	app.MountThreadsController(srv, threadCtrl)
	app.MountContentController(srv, contentCtrl)

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
	NewRoomController, wire.Bind(new(app.RoomsController), new(*RoomController)),
	NewUserRoomController, wire.Bind(new(app.UserRoomsController), new(*UserRoomController)),
	NewThreadController, wire.Bind(new(app.ThreadsController), new(*ThreadController)),
	NewContentController, wire.Bind(new(app.ContentController), new(*ContentController)),
)
