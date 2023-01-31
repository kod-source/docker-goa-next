package main

import (
	"database/sql"
	"errors"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
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
	myID := getUserIDCode(ctx)
	tu, err := t.tu.Create(ctx, pl.Text, model.RoomID(pl.RoomID), model.UserID(myID), pl.Img)
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

	return ctx.Created(toAppThreadUser(tu))
}

// Delete スレッドの削除
func (t *ThreadController) Delete(ctx *app.DeleteThreadsContext) error {
	if err := t.tu.Delete(ctx, model.UserID(getUserIDCode(ctx)), model.ThreadID(ctx.ID)); err != nil {
		switch err {
		case sql.ErrNoRows:
			return ctx.NotFound()
		case myerrors.ErrBadRequestNoPermission:
			return ctx.BadRequest()
		default:
			return ctx.InternalServerError()
		}
	}
	return ctx.OK(nil)
}

// GetThreadsByRoom ルーム内のスレッドを返す
func (t *ThreadController) GetThreadsByRoom(ctx *app.GetThreadsByRoomThreadsContext) error {
	its, nextID, err := t.tu.GetThreadsByRoom(ctx, model.RoomID(ctx.ID), model.ThreadID(pointer.Value(ctx.NextID)))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(toAllIndexThreads(its, nextID))
}

func toAppThreadUser(tu *model.ThreadUser) *app.ThreadUser {
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

func toAllIndexThreads(its []*model.IndexThread, nextID *int) *app.AllIndexThreads {
	var aits []*app.IndexThread
	for _, it := range its {
		aits = append(aits, &app.IndexThread{
			CountContent: it.CountContent,
			ThreadUser:   toAppThreadUser(&it.ThreadUser),
		})
	}

	return &app.AllIndexThreads{
		IndexThreads: aits,
		NextID:       nextID,
	}
}
