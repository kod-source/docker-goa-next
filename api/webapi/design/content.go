package design

import (
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
})
