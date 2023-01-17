package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// ContentController ...
type ContentController struct {
	*goa.Controller
}

// NewContentController ...
func NewContentController(service *goa.Service) *ContentController {
	return &ContentController{Controller: service.NewController("ThreadController")}
}

// Delete コンテントを削除する
func (c *ContentController) Delete(ctx *app.DeleteContentContext) error {
	return ctx.OK(nil)
}
