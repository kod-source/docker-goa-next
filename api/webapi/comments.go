package main

import (
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// CommentsController implements the comments resource.
type CommentsController struct {
	*goa.Controller
	cu usecase.CommentUsecase
}

// NewCommentsController creates a comments controller.
func NewCommentsController(service *goa.Service, cu usecase.CommentUsecase) *CommentsController {
	return &CommentsController{Controller: service.NewController("CommentsController"), cu: cu}
}

// CreateComment runs the create_comment action.
func (c *CommentsController) CreateComment(ctx *app.CreateCommentCommentsContext) error {
	comment, err := c.cu.Create(ctx, ctx.Payload.Text, ctx.Payload.Img)
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.Created(&app.CommentJSON{
		ID:        *comment.ID,
		PostID:    *comment.PostID,
		Text:      *comment.Text,
		Img:       comment.Img,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	})
}
