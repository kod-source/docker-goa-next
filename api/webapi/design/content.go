package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("content", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	BasePath("/content")

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Description("コンテントの削除")
		Params(func() {
			Param("id", Integer, "content id")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("create", func() {
		Routing(POST("/"))
		Description("スレッドの返信の作成")
		Payload(func() {
			Attribute("text", String, "コンテントの内容", func() {
				Example("いいですね")
			})
			Attribute("thread_id", Integer, "スレッドID", func() {
				Example(1)
			})
			Attribute("img", String, "画像", func() {
				Example("https://test.img")
			})
			Required("text", "thread_id")
		})
		Response(Created, contentUser)
		Response(BadRequest)
		Response(InternalServerError)
	})
})

var content = MediaType("application/vnd.content", func() {
	Description("コンテント")
	Attribute("id", Integer, "content id")
	Attribute("user_id", Integer, "user id")
	Attribute("thread_id", Integer, "thread id")
	Attribute("text", String, "コンテント内容")
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
		Attribute("thread_id")
		Attribute("text")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("img")
	})
	Required("id", "user_id", "thread_id", "text", "created_at", "updated_at")
})

var contentUser = MediaType("application/vnd.content_user", func() {
	Description("コンテントとユーザー")
	Attribute("content", content)
	Attribute("user", show_user)
	View("default", func() {
		Attribute("content")
		Attribute("user")
	})
	Required("content", "user")
})
