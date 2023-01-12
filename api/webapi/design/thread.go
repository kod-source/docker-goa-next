package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("threads", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("create", func() {
		Routing(POST("threads"))
		Description("スレッドの作成")
		Payload(func() {
			Attribute("text", String, "スレッドの内容", func() {
				Example("テストスレ")
			})
			Attribute("room_id", Integer, "ルームID", func() {
				Example(1)
			})
			Attribute("user_id", Integer, "ユーザーID", func() {
				Example(2)
			})
			Attribute("img", String, "画像", func() {
				Example("https://test.img")
			})
			Required("text", "room_id", "user_id")
		})
		Response(Created, threadUser)
		Response(BadRequest)
		Response(InternalServerError)
	})
})

var thread = MediaType("application/vnd.thread", func() {
	Description("スレッド")
	Attribute("id", Integer, "room id")
	Attribute("user_id", Integer, "user id")
	Attribute("room_id", Integer, "room id")
	Attribute("text", String, "スレッド内容")
	Attribute("created_at", DateTime, "作成日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("img", String, "画像")
	View("default", func() {
		Attribute("id")
		Attribute("user_id")
		Attribute("room_id")
		Attribute("text")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("img")
	})
	Required("id", "user_id", "room_id", "text", "created_at", "updated_at")
})

var threadUser = MediaType("application/vnd.thread_user", func() {
	Description("スレッドとユーザー情報")
	Attribute("thread", thread, "スレッド")
	Attribute("user", show_user, "ユーザー")
	View("default", func() {
		Attribute("thread")
		Attribute("user")
	})
	Required("thread", "user")
})
