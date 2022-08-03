package main

import (
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	uu usecase.UserUseCase
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, uu usecase.UserUseCase) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController"), uu: uu}
}

// GetCurrentUser runs the get_current_user action.
func (c *UsersController) GetCurrentUser(ctx *app.GetCurrentUserUsersContext) error {
	id := getUserIDCode(ctx)
	user, err := c.uu.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.NotFound()
		}
		return ctx.InternalServerError()
	}
	res := &app.User{
		ID:        user.ID,
		Email:     &user.Email,
		Name:      &user.Name,
		Password:  &user.Password,
		CreatedAt: &user.CreatedAt,
	}
	return ctx.OK(res)
}

func (c *UsersController) SignUp(ctx *app.SignUpUsersContext) error {
	user, err := c.uu.SignUp(ctx, ctx.Payload.Name, ctx.Payload.Email, ctx.Payload.Password)
	if err != nil {
		return ctx.BadRequest()
	}
	token, err := c.uu.CreateJWTToken(ctx, user.ID, user.Name)
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.Created(&app.Token{
		Token: *token,
		User: &app.User{
			CreatedAt: &user.CreatedAt,
			Email:     &user.Email,
			ID:        user.ID,
			Name:      &user.Name,
			Password:  &user.Password,
		},
	})
}
