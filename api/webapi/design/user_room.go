package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("user_rooms", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("invite_room", func() {
		Routing(POST("user_room"))
		Description("ルームに招待する")
		Payload(func() {
			Attribute("room_id", Integer, "ルームID", func() {
				Example(2)
			})
			Attribute("user_id", Integer, "ユーザーID", func() {
				Example(3)
			})
			Required("room_id", "user_id")
		})
		Response(Created, userRoom)
		Response(BadRequest)
		Response(InternalServerError)
	})

	Action("delete", func() {
		Routing(DELETE("user_room/:id"))
		Description("ルームから除外する")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK)
		Response(InternalServerError)
	})
})

var userRoom = MediaType("application/vnd.user_room+json", func() {
	Description("user room")
	Attribute("id", Integer, "ID", func() {
		Example(1)
	})
	Attribute("room_id", Integer, "ルームID", func() {
		Example(2)
	})
	Attribute("user_id", Integer, "ユーザーID", func() {
		Example(3)
	})
	Attribute("last_read_at", DateTime, "最後に開いた日時", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("created_at", DateTime, "作成日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	View("default", func() {
		Attribute("id")
		Attribute("room_id")
		Attribute("user_id")
		Attribute("last_read_at")
		Attribute("created_at")
		Attribute("updated_at")
	})
	Required("id", "room_id", "user_id", "created_at", "updated_at")
})
