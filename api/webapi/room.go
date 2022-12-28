package main

import (
	"database/sql"
	"errors"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
)

// RoomController ...
type RoomController struct {
	*goa.Controller
	ru usecase.RoomUseCase
}

// NewRoomController ...
func NewRoomController(service *goa.Service, ru usecase.RoomUseCase) *RoomController {
	return &RoomController{Controller: service.NewController("RoomController"), ru: ru}
}

// CreateRoom ルーム作成
func (r *RoomController) CreateRoom(ctx *app.CreateRoomRoomsContext) error {
	ru, err := r.ru.Create(ctx, ctx.Payload.Name, ctx.Payload.IsGroup, r.toUserIDsArray(ctx.Payload.UserIds))
	if err != nil {
		if errors.Is(err, myerrors.ErrBadRequestEmptyArray) {
			return ctx.BadRequest(&app.ServiceVerror{
				Code:    400,
				Message: "ユーザーIDを指定してください",
				Status:  err.Error(),
			})
		}
		return err
	}

	return ctx.Created(&app.IndexRooUser{
		ID:        int(ru.Room.ID),
		IsGroup:   ru.Room.IsGroup,
		Name:      ru.Room.Name,
		CreatedAt: ru.Room.CreatedAt,
		UpdatedAt: ru.Room.UpdatedAt,
		Users:     r.toShowUserCollection(ru.Users),
	})
}

func (r *RoomController) Index(ctx *app.IndexRoomsContext) error {
	rus, nextID, err := r.ru.Index(ctx, model.UserID(getUserIDCode(ctx)), model.RoomID(pointer.Value(ctx.NextID)))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	return ctx.OK(r.toAllRommUser(rus, nextID))
}

func (r *RoomController) toAllRommUser(rus []*model.RoomUser, nextID *int) *app.AllRoomUser {
	var iru []*app.IndexRooUser
	for _, ru := range rus {
		iru = append(iru, r.toRoomUser(ru))
	}

	return &app.AllRoomUser{
		NextID: nextID,
		Rooms:  iru,
	}
}

func (r *RoomController) toRoomUser(ru *model.RoomUser) *app.IndexRooUser {
	return &app.IndexRooUser{
		ID:        int(ru.Room.ID),
		IsGroup:   ru.Room.IsGroup,
		Name:      ru.Room.Name,
		CreatedAt: ru.Room.CreatedAt,
		UpdatedAt: ru.Room.UpdatedAt,
		Users:     r.toShowUserCollection(ru.Users),
	}
}

func (r *RoomController) toShowUserCollection(showUsers []*model.ShowUser) []*app.ShowUser {
	var sus []*app.ShowUser
	for _, su := range showUsers {
		sus = append(sus, &app.ShowUser{
			ID:        int(su.ID),
			Name:      su.Name,
			CreatedAt: su.CreatedAt,
			Avatar:    su.Avatar,
		})
	}

	return sus
}

func (r *RoomController) toUserIDsArray(ids []int) []model.UserID {
	var userIDs []model.UserID
	for _, id := range ids {
		userIDs = append(userIDs, model.UserID(id))
	}

	return userIDs
}
