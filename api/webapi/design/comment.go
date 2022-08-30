package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("comments", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("create_comment", func() {
		Routing(POST("comments"))
		Description("コメントの作成")
		Payload(func() {
			Attribute("post_id", Integer, "投稿ID", func() {
				Example(1)
			})
			Attribute("text", String, "コメントの内容", func() {
				Example("どうも~")
			})
			Attribute("img", String, "コメント画像のパス", func() {
				Example("data:image/jpeg;base64,/9j/4A")
			})
			Required("post_id", "text")
		})
		Response(Created, comment_with_user)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("show_comment", func() {
		Routing(GET("comments/:id"))
		Description("投稿に紐づくコメントの取得")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK, CollectionOf(comment_with_user))
		Response(NotFound)
		Response(InternalServerError)
	})

	Action("update_comment", func() {
		Routing(PUT("comments/:id"))
		Description("コメントの更新")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Payload(func() {
			Attribute("text", String, "コメントの内容", func() {
				Example("どうも~")
			})
			Attribute("img", String, "コメント画像のパス", func() {
				Example("data:image/jpeg;base64,/9j/4A")
			})
			Required("text")
		})
		Response(OK, comment)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("delete_comment", func() {
		Routing(DELETE("comments/:id"))
		Description("コメントの削除")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK)
		Response(NotFound)
		Response(InternalServerError)
	})
})

var comment = MediaType("application/vnd.comment_json", func() {
	Description("コメント")
	Attribute("id", Integer, "ID", func() {
		Example(1)
	})
	Attribute("post_id", Integer, "投稿ID", func() {
		Example(1)
	})
	Attribute("user_id", Integer, "ユーザーID", func() {
		Example(1)
	})
	Attribute("text", String, "コメント", func() {
		Example("やっほー")
	})
	Attribute("img", String, "コメントの画像パス", func() {
		Example("data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD/4QF8RXhpZgAATU0AKgAAAAgABgESAAMAAAABAAEAAAEaAAUAAAABAAAAVgEbAAUAAAABAAAAXgEoAAMAAAABAAIAAAEyAAIAAAAUAAAAZodpAAQAAAABAAAAegAAAAAAAABIAAAAAQAAAEgAAAABMjAyMjowNjoxOCAwMzo1NDo0MwAAD5AAAAcAAAAEMDIyMZADAAIAAAAUAAABNJAEAAIAAAAUAAABSJAQAAIAAAAHAAABXJARAAIAAAAHAAABZJASAAIAAAAHAAABbJEBAAcAAAAEAQIDAJKQAAIAAAAEOTIyAJKRAAIAAAAEOTIyAJKSAAIAAAAEOTIyAKAAAAcAAAAEMDEwMKABAAMAAAABAAEAAKACAAQAAAABAAACWKADAAQAAAABAAACWKQGAAMAAAABAAAAAAAAAAAyMDIyOjA2OjE4IDAzOjU0OjQzADIwMjI6MDY6MTggMDM6NTQ6NDMAKzA5OjAwAAArMDk6MDAAACswOTowMAAA/+0AeFBob3Rvc2hvcCAzLjAAOEJJTQQEAAAAAAA/HAFaAAMbJUccAgAAAgACHAI/AAYwMzU0NDMcAj4ACDIwMjIwNjE4HAI3AAgyMDIyMDYxOBwCPAAGMDM1NDQzADhCSU0EJQAAAAAAEKnEz4ubluvj5vP007FySPv/wAARCAJYAlgDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/")
	})
	Attribute("created_at", DateTime, "作成日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	View("default", func() {
		Attribute("id")
		Attribute("post_id")
		Attribute("user_id")
		Attribute("text")
		Attribute("img")
		Attribute("created_at")
		Attribute("updated_at")
	})
	Required("id", "post_id", "user_id", "text")
})

var comment_with_user = MediaType("application/vnd.comment_with_user_json", func() {
	Description("コメントとユーザーの情報")
	Attribute("comment", comment, "コメント")
	Attribute("user", user, "ユーザー")
	View("default", func() {
		Attribute("comment")
		Attribute("user")
	})
	Required("comment", "user")
})
