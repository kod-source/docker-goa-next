package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// ThreadController ...
type ThreadController struct {
	*goa.Controller
}

// NewThreadController ...
func NewThreadController(service *goa.Service) *ThreadController {
	return &ThreadController{Controller: service.NewController("ThreadController")}
}

// Create スレッド作成
func (t *ThreadController) Create(ctx *app.CreateThreadsContext) error {
	return ctx.Created(nil)
}
