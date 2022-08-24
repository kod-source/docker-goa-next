package main

import (
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// LikesController implements the likes resource.
type LikesController struct {
	*goa.Controller
	lu usecase.LikeUsecase
}

// NewLikesController creates a likes controller.
func NewLikesController(service *goa.Service, lu usecase.LikeUsecase) *LikesController {
	return &LikesController{Controller: service.NewController("LikesController"), lu: lu}
}

// Create runs the create action.
func (c *LikesController) Create(ctx *app.CreateLikesContext) error {
	l, err := c.lu.Create(ctx, getUserIDCode(ctx), ctx.Payload.PostID)
	if err != nil {
		if err == myerrors.BadRequestIntError {
			return ctx.BadRequest(&app.ServiceVerror{
				Code:    400,
				Message: "不明なリクエストです",
				Status:  err.Error(),
			})
		}
		if code := myerrors.GetMySQLErrorNumber(err); code == myerrors.MySQLErrorDuplicate.Number {
			return ctx.BadRequest(&app.ServiceVerror{
				Code:    int(code),
				Message: "ユニークエラーです",
				Status:  err.Error(),
			})
		}
		return ctx.InternalServerError()
	}

	return ctx.Created(&app.LikeJSON{
		ID:     l.ID,
		PostID: l.PostID,
		UserID: l.UserID,
	})
}

func (c *LikesController) Delete(ctx *app.DeleteLikesContext) error {
	err := c.lu.Delete(ctx, getUserIDCode(ctx), ctx.Payload.PostID)
	if err != nil {
		if err == myerrors.BadRequestIntError {
			return ctx.BadRequest()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(nil)
}

func (c *LikesController) GetMyLike(ctx *app.GetMyLikeLikesContext) error {
	postIDs := [...]int{1, 2, 3, 5}
	return ctx.OK(postIDs[:])
}
