package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// LikesController implements the likes resource.
type LikesController struct {
	*goa.Controller
}

// NewLikesController creates a likes controller.
func NewLikesController(service *goa.Service) *LikesController {
	return &LikesController{Controller: service.NewController("LikesController")}
}

// Create runs the create action.
func (c *LikesController) Create(ctx *app.CreateLikesContext) error {
	// LikesController_Create: start_implement

	// Put your logic here

	return nil
	// LikesController_Create: end_implement
}
