// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package webapi

import (
	"context"
	"github.com/kod-source/docker-goa-next/app/datastore"
	"github.com/kod-source/docker-goa-next/app/external"
	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/repository"
)

// Injectors from wire.go:

func NewApp(ctx context.Context) (*App, error) {
	service := newService()
	webapiAppConfig, err := getAppConfig()
	if err != nil {
		return nil, err
	}
	config, err := newMysqlConfig(webapiAppConfig)
	if err != nil {
		return nil, err
	}
	db, err := datastore.NewDB(config)
	if err != nil {
		return nil, err
	}
	timeRepository := repository.NewTimeRepositoy()
	userDatastore := datastore.NewUserDatastore(db, timeRepository)
	jwtExternal := external.NewJWTExternal(timeRepository)
	userInteractor := interactor.NewUserInteractor(userDatastore, jwtExternal)
	usersController := NewUsersController(service, userInteractor)
	postDatastore := datastore.NewPostDatastore(db, timeRepository)
	postInteractor := interactor.NewPostInteractor(postDatastore)
	postsController := NewPostsController(service, postInteractor)
	operandsController := NewOperandsController(service)
	likeDatastore := datastore.NewLikeDatastore(db)
	likeInteractor := interactor.NewLikeInteractor(likeDatastore)
	likesController := NewLikesController(service, likeInteractor)
	commentDatastore := datastore.NewCommentDatastore(db, timeRepository)
	commentInteractor := interactor.NewCommentInteractor(commentDatastore)
	commentsController := NewCommentsController(service, commentInteractor)
	oauth2Config, err := newGoogleConfig()
	if err != nil {
		return nil, err
	}
	googleExternal := external.NewGoogleService(oauth2Config)
	googleInteractor := interactor.NewGoogleUseCase(googleExternal, userDatastore)
	authController := NewAuthController(service, userInteractor, googleInteractor)
	roomDatastore := datastore.NewRoomDatastore(db, timeRepository)
	roomInteractor := interactor.NewRoomInterractor(roomDatastore)
	roomController := NewRoomController(service, roomInteractor)
	userRoomDatastore := datastore.NewUserRoomRepository(db, timeRepository)
	userRoomInteractor := interactor.NewUserRoomUsecase(userRoomDatastore)
	userRoomController := NewUserRoomController(service, userRoomInteractor)
	threadDatastore := datastore.NewThreadRepository(db, timeRepository)
	threadInteractor := interactor.NewThreadUsecase(threadDatastore)
	wsConnections := newConnections(ctx)
	threadController := NewThreadController(service, threadInteractor, wsConnections)
	contentDatastore := datastore.NewContentRepository(db, timeRepository)
	contentInteractor := interactor.NewContentUsecase(contentDatastore)
	contentController := NewContentController(service, contentInteractor)
	app, err := newApp(ctx, service, usersController, postsController, operandsController, likesController, commentsController, authController, roomController, userRoomController, threadController, contentController)
	if err != nil {
		return nil, err
	}
	return app, nil
}
