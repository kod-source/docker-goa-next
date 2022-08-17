package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// CommentsController implements the comments resource.
type CommentsController struct {
	*goa.Controller
}

// NewCommentsController creates a comments controller.
func NewCommentsController(service *goa.Service) *CommentsController {
	return &CommentsController{Controller: service.NewController("CommentsController")}
}

// CreateComment runs the create_comment action.
func (c *CommentsController) CreateComment(ctx *app.CreateCommentCommentsContext) error {
	// CommentsController_CreateComment: start_implement

	// Put your logic here

	return nil
	// CommentsController_CreateComment: end_implement
}
