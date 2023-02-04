package webapi

import (
	"context"
	"os"
	"testing"
	"time"
)

var testApp *App
var ctx context.Context
var jst *time.Location

func TestMain(m *testing.M) {
	testApp = &App{}
	testApp.srv = newService()
	ctx = context.Background()

	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Exit(1)
	}

	m.Run()
}
