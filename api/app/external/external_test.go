package external

import (
	"context"
	"os"
	"testing"
	"time"
)

var ctx context.Context
var jst *time.Location

func TestMain(m *testing.M) {
	ctx = context.Background()

	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Exit(1)
	}

	m.Run()
}
