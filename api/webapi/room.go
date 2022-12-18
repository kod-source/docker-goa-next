package main

import (
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// RoomController ...
type RoomController struct {
	*goa.Controller
}

// NewRoomController ...
func NewRoomController(service *goa.Service) *RoomController {
	return &RoomController{Controller: service.NewController("RoomController")}
}

func (r *RoomController) CreateRoom(ctx *app.CreateRoomRoomsContext) error {
	// usecaseの実装

	return ctx.Created(&app.IndexRooUser{})
}
