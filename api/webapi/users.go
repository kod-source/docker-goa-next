package main

import (
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	usecase.UserUseCase
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController"), UserUseCase: usecase.NewUserUseCase()}
}

// GetCurrentUser runs the get_current_user action.
func (c *UsersController) GetCurrentUser(ctx *app.GetCurrentUserUsersContext) error {
	id := getUserIDCode(ctx)
	user, err := c.GetUser(ctx, id)
	if err != nil {
		return ctx.NotFound()
	}
	res := &app.User{
		ID:       user.ID,
		Email:    &user.Email,
		Name:     &user.Name,
		Password: &user.Password,
	}
	return ctx.OK(res)
}
