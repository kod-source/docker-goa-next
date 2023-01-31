package main

import (
	"context"

	"github.com/kod-source/docker-goa-next/webapi"
)

func main() {
	app, err := webapi.NewApp(context.Background())
	if err != nil {
		panic(err)
	}

	// Start service
	if err := app.LaunchServer(); err != nil {
		panic(err)
	}
}
