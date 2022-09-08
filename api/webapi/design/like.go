package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("likes", func() {
	Security(JWT, func() {
		Scope("api:access")
	})

	Action("create", func() {
		Routing(POST("likes"))
		Description("いいね作成")
		Payload(func() {
			Attribute("post_id", Integer, "投稿ID", func() {
				Example(1)
			})
			Required("post_id")
		})
		Response(Created, like)
		Response(BadRequest, MyError)
		Response(InternalServerError)
	})

	Action("delete", func() {
		Routing(DELETE("likes"))
		Description("いいねの削除")
		Payload(func() {
			Attribute("post_id", Integer, "投稿ID", func() {
				Example(1)
			})
			Required("post_id")
		})
		Response(OK)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("get_my_like", func() {
		Routing(GET("likes"))
		Description("自分がいいねした全ての投稿IDを取得する")
		Response(OK, ArrayOf(Integer))
		Response(InternalServerError)
	})

	Action("get_like_by_user", func() {
		Routing(GET("likes/:user_id"))
		Params(func() {
			Param("user_id", Integer, "UserID")
		})
		Description("user_idに紐づくいいねした全ての投稿IDを取得する")
		Response(OK, ArrayOf(Integer))
		Response(NotFound)
		Response(InternalServerError)
	})
})

var like = MediaType("application/vnd.like_json", func() {
	Description("いいね")
	Attribute("id", Integer, "ID")
	Attribute("post_id", Integer, "投稿ID")
	Attribute("user_id", Integer, "ユーザーID")
	View("default", func() {
		Attribute("id")
		Attribute("post_id")
		Attribute("user_id")
	})
	Required("id", "post_id", "user_id")
})
