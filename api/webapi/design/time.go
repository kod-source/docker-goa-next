package design

import (
	"log"
	"time"
)

var loc *time.Location
var err error

func init() {
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Panic(err)
	}
}
