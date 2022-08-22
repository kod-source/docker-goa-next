package main

import (
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
		return ctx.InternalServerError()
	}

	return ctx.Created(&app.LikeJSON{
		ID:     l.ID,
		PostID: l.PostID,
		UserID: l.UserID,
	})
}
