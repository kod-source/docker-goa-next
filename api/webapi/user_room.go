package main

import (
	"errors"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
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
	userRoom, err := ur.uru.InviteRoom(ctx, model.RoomID(roomID), model.UserID(userID))
	if err != nil {
		if errors.Is(err, myerrors.ErrBadRequestInt) || errors.Is(err, myerrors.MySQLErrorAddOrUpdateForeignKey) {
			return ctx.BadRequest()
		}
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
