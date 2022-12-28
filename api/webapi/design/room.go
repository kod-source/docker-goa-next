package design

import (
	"time"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("rooms", func() {
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("create_room", func() {
		Routing(POST("rooms"))
		Description("ルームの作成")
		Payload(func() {
			Attribute("name", String, "ルーム名", func() {
				Example("DBルーム")
			})
			Attribute("is_group", Boolean, "DBかどうか", func() {
				Example(true)
			})
			Attribute("user_ids", ArrayOf(Integer), "ルームに入れるUserID", func() {
				Example([]int{1, 2})
			})
			Required("name", "is_group", "user_ids")
		})
		Response(Created, roomUser)
		Response(BadRequest, MyError)
		Response(InternalServerError)
	})

	Action("index", func() {
		Routing(GET("rooms"))
		Description("自分が入っている全てのルームを表示")
		Params(func() {
			Param("next_id", Integer, "次のID")
		})
		Response(OK, allRoomUser)
		Response(NotFound)
		Response(InternalServerError)
	})
})

var roomUser = MediaType("application/vnd.index_roo_user", func() {
	Description("ルーム")
	Attribute("id", Integer, "room id")
	Attribute("name", String, "room name")
	Attribute("is_group", Boolean, "グループかDMの判定")
	Attribute("created_at", DateTime, "作成日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("updated_at", DateTime, "更新日", func() {
		Example(time.Date(2019, 01, 31, 0, 0, 0, 0, loc).Format(time.RFC3339))
	})
	Attribute("users", CollectionOf(show_user), "ルームいるユーザー")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("is_group")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("users")
	})
	Required("id", "name", "is_group", "created_at", "updated_at", "users")
})

var allRoomUser = MediaType("application/vnd.all_room_user", func() {
	Description("全てのルーム")
	Attribute("rooms", CollectionOf(roomUser), "rooms")
	Attribute("next_id", Integer, "次取得するRoomID")
	View("default", func() {
		Attribute("rooms")
		Attribute("next_id")
	})
	Required("rooms")
})
