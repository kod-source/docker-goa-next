package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("users", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("get_current_user", func() {
		Routing(GET("current_user"))
		Description("ログインしているユーザーの情報を取得する")
		Response(OK, user)
		Response(NotFound)
		Response(InternalServerError)
	})
})

var user = MediaType("application/vnd.user+json", func() {
	Description("user")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "name", func() {
		Example("佐藤　太郎")
	})
	Attribute("email", String, "email", func() {
		Example("test@exmaple.com")
	})
	Attribute("password", String, "password", func() {
		Example("pas")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("password")
	})
	Required("id")
})
