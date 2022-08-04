package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// PostsController implements the posts resource.
type PostsController struct {
	*goa.Controller
}

// NewPostsController creates a posts controller.
func NewPostsController(service *goa.Service) *PostsController {
	return &PostsController{Controller: service.NewController("PostsController")}
}

// CreatePost runs the create_post action.
func (c *PostsController) CreatePost(ctx *app.CreatePostPostsContext) error {
	// PostsController_CreatePost: start_implement

	// Put your logic here

	return nil
	// PostsController_CreatePost: end_implement
}
