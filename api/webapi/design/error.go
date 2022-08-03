package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var MyError = MediaType("application/vnd.service.verror", func() {
	Description("my error")
	Attributes(func() {
		Attribute("code", Integer, "Code", func() {
			Example(404)
		})
		Attribute("details", Any, "Details")
		Attribute("status", String, "Status", func() {
			Example("unauthenticated")
		})
		Attribute("message", String, "エラーメッセージ", func() {
			Example("record not found")
		})
		Required("code", "status", "message")
	})
	View("default", func() {
		Attribute("code")
		Attribute("details")
		Attribute("status")
		Attribute("message")
	})
})
