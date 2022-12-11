//go:generate go run gen/main.go

package main

import (
	"context"
)

func main() {
	app, err := NewApp(context.Background())
	if err != nil {
		panic(err)
	}

	// Start service
	if err := app.srv.ListenAndServe(":3000"); err != nil {
		app.srv.LogError("startup", "err", err)
	}
}
