package main

import (
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
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
	comment, err := c.cu.Create(ctx, ctx.Payload.PostID, ctx.Payload.Text, ctx.Payload.Img)
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

func (c *CommentsController) ShowComment(ctx *app.ShowCommentCommentsContext) error {
	cs, err := c.cu.ShowByPostID(ctx, ctx.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(c.ToCommentJSONCollection(cs))
}

func (c *CommentsController) ToCommentJSONCollection(comments []*model.Comment) app.CommentJSONCollection {
	var cs app.CommentJSONCollection
	for _, c := range comments {
		cs = append(cs, &app.CommentJSON{
			ID:        *c.ID,
			PostID:    *c.PostID,
			Img:       c.Img,
			Text:      *c.Text,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return cs
}
