package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("posts", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("create_post", func() {
		Routing(POST("posts"))
		Description("投稿を作成する")
		Payload(func() {
			Attribute("title", String, "タイトル", func() {
				Example("やっほー")
			})
			Attribute("img", String, "プロフィール画像のパス", func() {
				Example("data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD/4QF8RXhpZgAATU0AKgAAAAgABgESAAMAAAABAAEAAAEaAAUAAAABAAAAVgEbAAUAAAABAAAAXgEoAAMAAAABAAIAAAEyAAIAAAAUAAAAZodpAAQAAAABAAAAegAAAAAAAABIAAAAAQAAAEgAAAABMjAyMjowNjoxOCAwMzo1NDo0MwAAD5AAAAcAAAAEMDIyMZADAAIAAAAUAAABNJAEAAIAAAAUAAABSJAQAAIAAAAHAAABXJARAAIAAAAHAAABZJASAAIAAAAHAAABbJEBAAcAAAAEAQIDAJKQAAIAAAAEOTIyAJKRAAIAAAAEOTIyAJKSAAIAAAAEOTIyAKAAAAcAAAAEMDEwMKABAAMAAAABAAEAAKACAAQAAAABAAACWKADAAQAAAABAAACWKQGAAMAAAABAAAAAAAAAAAyMDIyOjA2OjE4IDAzOjU0OjQzADIwMjI6MDY6MTggMDM6NTQ6NDMAKzA5OjAwAAArMDk6MDAAACswOTowMAAA/+0AeFBob3Rvc2hvcCAzLjAAOEJJTQQEAAAAAAA/HAFaAAMbJUccAgAAAgACHAI/AAYwMzU0NDMcAj4ACDIwMjIwNjE4HAI3AAgyMDIyMDYxOBwCPAAGMDM1NDQzADhCSU0EJQAAAAAAEKnEz4ubluvj5vP007FySPv/wAARCAJYAlgDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/")
			})
			Required("title")
		})
		Response(Created, post_and_user)
		Response(BadRequest, MyError)
		Response(InternalServerError)
	})

	Action("index", func() {
		Routing(GET("posts"))
		Params(func() {
			Param("next_id", Integer, "次のID")
		})
		Description("全部の登録を取得する")
		Response(OK, post_all_limit)
		Response(NotFound)
		Response(InternalServerError)
	})

	Action("delete", func() {
		Routing(DELETE("posts/:id"))
		Description("投稿を削除する")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK)
		Response(InternalServerError)
	})

	Action("update", func() {
		Routing(PUT("posts/:id"))
		Description("投稿を更新する")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Payload(func() {
			Attribute("title", String, "タイトル", func() {
				Example("やっほー")
			})
			Attribute("img", String, "プロフィール画像のパス", func() {
				Example("data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD/4QF8RXhpZgAATU0AKgAAAAgABgESAAMAAAABAAEAAAEaAAUAAAABAAAAVgEbAAUAAAABAAAAXgEoAAMAAAABAAIAAAEyAAIAAAAUAAAAZodpAAQAAAABAAAAegAAAAAAAABIAAAAAQAAAEgAAAABMjAyMjowNjoxOCAwMzo1NDo0MwAAD5AAAAcAAAAEMDIyMZADAAIAAAAUAAABNJAEAAIAAAAUAAABSJAQAAIAAAAHAAABXJARAAIAAAAHAAABZJASAAIAAAAHAAABbJEBAAcAAAAEAQIDAJKQAAIAAAAEOTIyAJKRAAIAAAAEOTIyAJKSAAIAAAAEOTIyAKAAAAcAAAAEMDEwMKABAAMAAAABAAEAAKACAAQAAAABAAACWKADAAQAAAABAAACWKQGAAMAAAABAAAAAAAAAAAyMDIyOjA2OjE4IDAzOjU0OjQzADIwMjI6MDY6MTggMDM6NTQ6NDMAKzA5OjAwAAArMDk6MDAAACswOTowMAAA/+0AeFBob3Rvc2hvcCAzLjAAOEJJTQQEAAAAAAA/HAFaAAMbJUccAgAAAgACHAI/AAYwMzU0NDMcAj4ACDIwMjIwNjE4HAI3AAgyMDIyMDYxOBwCPAAGMDM1NDQzADhCSU0EJQAAAAAAEKnEz4ubluvj5vP007FySPv/wAARCAJYAlgDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/")
			})
			Required("title")
		})
		Response(OK, post_and_user)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("show", func() {
		Routing(GET("posts/:id"))
		Description("一つの投稿を取得する")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK, show_post)
		Response(NotFound)
		Response(InternalServerError)
	})

	Action("show_my_like", func() {
		Routing(GET("posts/my_like"))
		Description("自分がいいねした投稿を取得する")
		Params(func() {
			Param("next_id", Integer, "次のID")
		})
		Response(OK, post_all_limit)
		Response(NotFound)
		Response(InternalServerError)
	})

	Action("show_post_like", func() {
		Routing(GET("posts/likes/:id"))
		Description("指定したユーザーIDのいいねした投稿を取得する")
		Params(func() {
			Param("id", Integer, "User ID")
			Param("next_id", Integer, "次のID")
		})
		Response(OK, post_all_limit)
		Response(NotFound)
		Response(InternalServerError)
	})
})

var post = MediaType("application/vnd.post_json", func() {
	Description("投稿")
	Attribute("id", Integer, "ID", func() {
		Example(1)
	})
	Attribute("user_id", Integer, "ユーザーID", func() {
		Example(1)
	})
	Attribute("title", String, "タイトル", func() {
		Example("やっほー")
	})
	Attribute("img", String, "プロフィール画像のパス", func() {
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
		Attribute("user_id")
		Attribute("title")
		Attribute("img")
		Attribute("created_at")
		Attribute("updated_at")
	})
	Required("id", "user_id", "title")
})

var post_and_user = MediaType("application/vnd.index_post_json", func() {
	Description("投稿")
	Attribute("post", post, "post value")
	Attribute("user_name", String, "ユーザー名", func() {
		Example("佐藤　太郎")
	})
	Attribute("avatar", String, "ユーザー名", func() {
		Example("data:image/jpg")
	})
	View("default", func() {
		Attribute("post")
		Attribute("user_name")
		Attribute("avatar")
	})
	Required("post", "user_name")
})

var post_and_all_user = MediaType("application/vnd.post_and_user_json", func() {
	Description("投稿とユーザーの情報")
	Attribute("post", post, "post value")
	Attribute("user", user, "user value")
	View("default", func() {
		Attribute("post")
		Attribute("user")
	})
	Required("post", "user")
})

var show_post = MediaType("application/vnd.show_post_json", func() {
	Description("投稿とユーザーとコメントの情報")
	Attribute("post", post, "post value")
	Attribute("user", user, "user value")
	Attribute("comments_with_users", CollectionOf(comment_with_user), "comments value")
	Attribute("likes", CollectionOf(like), "likes value")
	View("default", func() {
		Attribute("post")
		Attribute("user")
		Attribute("comments_with_users")
		Attribute("likes")
	})
	Required("post", "user", "comments_with_users", "likes")
})

var post_all_limit = MediaType("application/vnd.post_all_limit", func() {
	Description("投稿とnext_idに情報")
	Attribute("show_posts", CollectionOf(post_and_user_and_count_like), "post_and_user vbalue")
	Attribute("next_token", String, "http://localhost:3000/posts?next_id=20")
	View("default", func() {
		Attribute("show_posts")
		Attribute("next_token")
	})
	Required("show_posts")
})

var post_and_user_and_count_like = MediaType("application/vnd.post_and_user_and_count_like_json", func() {
	Description("投稿といいね数")
	Attribute("post", post, "post value")
	Attribute("user_name", String, "ユーザー名", func() {
		Example("佐藤　太郎")
	})
	Attribute("avatar", String, "ユーザー名", func() {
		Example("data:image/jpg")
	})
	Attribute("count_like", Integer, "いいね数", func() {
		Example(10)
	})
	Attribute("count_comment", Integer, "コメント数", func() {
		Example(10)
	})
	View("default", func() {
		Attribute("post")
		Attribute("user_name")
		Attribute("avatar")
		Attribute("count_like")
		Attribute("count_comment")
	})
	Required("post", "user_name", "count_like", "count_comment")
})
