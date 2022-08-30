package main

import (
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
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
	cu, err := c.cu.Create(ctx, ctx.Payload.PostID, getUserIDCode(ctx), ctx.Payload.Text, ctx.Payload.Img)
	if err != nil {
		if err == myerrors.BadRequestStingError {
			return ctx.BadRequest()
		}
		return ctx.InternalServerError()
	}
	return ctx.Created(&app.CommentWithUserJSON{
		Comment: &app.CommentJSON{
			ID:        pointer.IntValue(cu.Comment.ID),
			PostID:    pointer.IntValue(cu.Comment.PostID),
			UserID:    pointer.IntValue(cu.Comment.UserID),
			Img:       cu.Comment.Img,
			Text:      pointer.StringValue(cu.Comment.Text),
			CreatedAt: cu.Comment.CreatedAt,
			UpdatedAt: cu.Comment.UpdatedAt,
		},
		User: &app.User{
			ID:     cu.User.ID,
			Name:   &cu.User.Name,
			Avatar: cu.User.Avatar,
		},
	})
}

func (c *CommentsController) ShowComment(ctx *app.ShowCommentCommentsContext) error {
	cs, err := c.cu.ShowByPostID(ctx, ctx.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}

	return ctx.OK(c.ToCommentJSONCollection(cs))
}

func (c *CommentsController) UpdateComment(ctx *app.UpdateCommentCommentsContext) error {
	comment, err := c.cu.Update(ctx, ctx.ID, ctx.Payload.Text, ctx.Payload.Img)
	if err != nil {
		if err == myerrors.BadRequestStingError {
			return ctx.BadRequest()
		}
		return ctx.InternalServerError()
	}
	return ctx.OK(&app.CommentJSON{
		CreatedAt: comment.CreatedAt,
		ID:        *comment.ID,
		Img:       comment.Img,
		PostID:    *comment.PostID,
		UserID:    *comment.UserID,
		Text:      *comment.Text,
		UpdatedAt: comment.UpdatedAt,
	})
}

func (c *CommentsController) DeleteComment(ctx *app.DeleteCommentCommentsContext) error {
	err := c.cu.Delete(ctx, ctx.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	return ctx.OK(nil)
}

func (c *CommentsController) ToCommentJSONCollection(comments_with_user []*model.CommentWithUser) app.CommentWithUserJSONCollection {
	cus := make(app.CommentWithUserJSONCollection, 0, len(comments_with_user))
	for _, cu := range comments_with_user {
		cus = append(cus, &app.CommentWithUserJSON{
			Comment: &app.CommentJSON{
				ID:        pointer.IntValue(cu.Comment.ID),
				PostID:    pointer.IntValue(cu.Comment.PostID),
				UserID:    pointer.IntValue(cu.Comment.UserID),
				Text:      pointer.StringValue(cu.Comment.Text),
				Img:       cu.Comment.Img,
				CreatedAt: cu.Comment.CreatedAt,
				UpdatedAt: cu.Comment.UpdatedAt,
			},
			User: &app.User{
				ID:        cu.User.ID,
				Name:      &cu.User.Name,
				Email:     &cu.User.Email,
				Avatar:    cu.User.Avatar,
				CreatedAt: &cu.User.CreatedAt,
			},
		})
	}
	return cus
}
