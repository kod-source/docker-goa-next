//go:generate goagen bootstrap -d github.com/kod-source/docker-goa-next/webapi/design

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	// Create service
	service := goa.New("docker_goa_next")
	fmt.Println("実行")
	_, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
	}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.UseJWTMiddleware(service, newAuthMiddleware())

	// Mount "operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	a := NewAuthController(service)
	app.MountAuthController(service, a)
	u := NewUsersController(service)
	app.MountUsersController(service, u)

	// Start service
	if err := service.ListenAndServe(":3000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
