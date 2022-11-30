//go:generate go run gen/main.go

package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// // Create service
	// service := goa.New("docker_goa_next")

	// // Mount middleware
	// service.Use(middleware.RequestID())
	// service.Use(middleware.LogRequest(true))
	// service.Use(middleware.ErrorHandler(service, true))
	// service.Use(middleware.Recover())

	// app.UseJWTMiddleware(service, newAuthMiddleware())

	// db, err := schema.NewDB()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer db.Close()
	// // Mount "operands" controller
	// uu := usecase.NewUserUseCase(datastore.NewUserDatastore(db, repository.NewTimeRepositoy()), datastore.NewJWTDatastore(repository.NewTimeRepositoy()))
	// pu := usecase.NewPostUseCase(datastore.NewPostDatastore(db, repository.NewTimeRepositoy()))
	// c := NewOperandsController(service)
	// app.MountOperandsController(service, c)
	// a := NewAuthController(service, uu)
	// app.MountAuthController(service, a)
	// u := NewUsersController(service, uu)
	// app.MountUsersController(service, u)
	// p := NewPostsController(service, pu)
	// app.MountPostsController(service, p)
	// cc := NewCommentsController(service, usecase.NewCommentUsecase(datastore.NewCommentDatastore(db, repository.NewTimeRepositoy())))
	// app.MountCommentsController(service, cc)
	// l := NewLikesController(service, usecase.NewLikeUsecase(datastore.NewLikeDatastore(db)))
	// app.MountLikesController(service, l)

	// // Start service
	// if err := service.ListenAndServe(":3000"); err != nil {
	// 	service.LogError("startup", "err", err)
	// }

	app, err := NewApp(context.Background())
	if err != nil {
		panic(err)
	}
	if err := app.srv.ListenAndServe(":3000"); err != nil {
		app.srv.LogError("startup", "err", err)
	}
}
