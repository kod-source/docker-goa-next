package webapi

import (
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// ContentController ...
type ContentController struct {
	*goa.Controller
	cu usecase.ContentUsecase
}

// NewContentController ...
func NewContentController(service *goa.Service, cu usecase.ContentUsecase) *ContentController {
	return &ContentController{Controller: service.NewController("ThreadController"), cu: cu}
}

// Delete コンテントを削除する
func (c *ContentController) Delete(ctx *app.DeleteContentContext) error {
	if err := c.cu.Delete(ctx, model.UserID(getUserIDCode(ctx)), model.ContentID(ctx.ID)); err != nil {
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

// Create スレッドの返信を作成する
func (c *ContentController) Create(ctx *app.CreateContentContext) error {
	payload := ctx.Payload
	cu, err := c.cu.Create(ctx, payload.Text, model.ThreadID(payload.ThreadID), model.UserID(getUserIDCode(ctx)), payload.Img)
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

	return ctx.Created(toAppContentUser(cu))
}

// GetByThread スレッドの返信の一覧を返す
func (c *ContentController) GetByThread(ctx *app.GetByThreadContentContext) error {
	return ctx.OK(nil)
}

func toAppContentUser(cu *model.ContentUser) *app.ContentUser {
	return &app.ContentUser{
		Content: &app.Content{
			ID:        int(cu.Content.ID),
			Text:      cu.Content.Text,
			UserID:    int(cu.Content.UserID),
			ThreadID:  int(cu.Content.ThreadID),
			UpdatedAt: cu.Content.UpdatedAt,
			CreatedAt: cu.Content.CreatedAt,
			Img:       cu.Content.Img,
		},
		User: &app.ShowUser{
			ID:        int(cu.User.ID),
			Name:      cu.User.Name,
			CreatedAt: cu.User.CreatedAt,
			Avatar:    cu.User.Avatar,
		},
	}
}
