package main

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
	return ctx.Created(nil)
}
