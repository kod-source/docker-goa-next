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
	roomID := ctx.Payload.RoomID
	userID := ctx.Payload.UserID
	if roomID == 0 || userID == 0 {
		return ctx.BadRequest()
	}

	userRoom, err := ur.uru.InviteRoom(ctx, model.RoomID(roomID), model.UserID(userID))
	if err != nil {
		return ctx.InternalServerError()
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
