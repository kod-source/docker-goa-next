package main

import (
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// ThreadController ...
type ThreadController struct {
	*goa.Controller
	tu usecase.ThreadUsecase
}

// NewThreadController ...
func NewThreadController(service *goa.Service, tu usecase.ThreadUsecase) *ThreadController {
	return &ThreadController{Controller: service.NewController("ThreadController"), tu: tu}
}

// Create スレッド作成
func (t *ThreadController) Create(ctx *app.CreateThreadsContext) error {
	pl := ctx.Payload
	tu, err := t.tu.Create(ctx, pl.Text, model.RoomID(pl.RoomID), model.UserID(pl.UserID), pl.Img)
	if err != nil {
		if code := myerrors.GetMySQLErrorNumber(err); code == myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			return ctx.BadRequest()
		}
		switch err {
		case myerrors.ErrBadRequestSting, myerrors.ErrBadRequestInt:
			return ctx.BadRequest()
		default:
			return ctx.InternalServerError()
		}
	}

	return ctx.Created(t.toAppThreadUser(tu))
}

func (t *ThreadController) toAppThreadUser(tu *model.ThreadUser) *app.ThreadUser {
	return &app.ThreadUser{
		Thread: &app.Thread{
			ID:        int(tu.Thread.ID),
			UserID:    int(tu.Thread.UserID),
			RoomID:    int(tu.Thread.RoomID),
			Text:      tu.Thread.Text,
			CreatedAt: tu.Thread.CreatedAt,
			UpdatedAt: tu.Thread.UpdatedAt,
			Img:       tu.Thread.Img,
		},
		User: &app.ShowUser{
			ID:        int(tu.User.ID),
			Name:      tu.User.Name,
			CreatedAt: tu.User.CreatedAt,
			Avatar:    tu.User.Avatar,
		},
	}
}
