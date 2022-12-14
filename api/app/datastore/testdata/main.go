package testdata

import "time"

var jst *time.Location
var now time.Time

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	now = time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
}
