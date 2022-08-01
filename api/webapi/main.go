//go:generate goagen bootstrap -d github.com/kod-source/docker-goa-next/webapi/design

package main

import (
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	// Create service
	service := goa.New("docker_goa_next")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.UseJWTMiddleware(service, newAuthMiddleware())

	// Mount "operands" controller
	uu := usecase.NewUserUseCase()
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	a := NewAuthController(service, uu)
	app.MountAuthController(service, a)
	u := NewUsersController(service, uu)
	app.MountUsersController(service, u)

	// Start service
	if err := service.ListenAndServe(":3000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
