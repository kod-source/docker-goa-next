package main

import (
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// UserRoomController ...
type UserRoomController struct {
	*goa.Controller
	uru usecase.UserRoomUseCase
}

// NewUserRoomController ...
func NewUserRoomController(service *goa.Service, uru usecase.UserRoomUseCase) *UserRoomController {
	return &UserRoomController{Controller: service.NewController("UserRoomController"), uru: uru}
}

// InviteRoom ルームに招待する
func (ur *UserRoomController) InviteRoom(ctx *app.InviteRoomUserRoomsContext) error {
	userRoom, err := ur.uru.InviteRoom(ctx, model.RoomID(ctx.Payload.RoomID), model.UserID(ctx.Payload.UserID))
	if err != nil {
		ctx.InternalServerError()
	}
	return ctx.Created(&app.UserRoom{
		ID:         int(userRoom.ID),
		RoomID:     int(userRoom.RoomID),
		UserID:     int(userRoom.UserID),
		LastReadAt: userRoom.LastReadAt,
		CreatedAt:  userRoom.CreatedAt,
		UpdatedAt:  userRoom.UpdatedAt,
	})
}
