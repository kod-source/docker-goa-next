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
			Attribute("text", String, "コメントの内容", func() {
				Example("どうも~")
			})
			Attribute("img", String, "コメント画像のパス", func() {
				Example("data:image/jpeg;base64,/9j/4A")
			})
			Required("text")
		})
		Response(Created, comment)
		Response(BadRequest)
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
		Attribute("text")
		Attribute("img")
		Attribute("created_at")
		Attribute("updated_at")
	})
	Required("id", "post_id", "text")
})
