package design

import (
	"time"

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
	Attribute("id", Integer, "ID", func() {
		Example(1)
	})
	Attribute("name", String, "名前", func() {
		Example("佐藤　太郎")
	})
	Attribute("email", String, "メール", func() {
		Example("test@exmaple.com")
	})
	Attribute("password", String, "パスワード", func() {
		Example("pas")
	})
	Attribute("created_at", DateTime, "作成日", func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("password")
		Attribute("created_at")
	})
	Required("id")
})
