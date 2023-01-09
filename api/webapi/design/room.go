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
				Example("DMルーム")
			})
			Attribute("is_group", Boolean, "DMかどうか", func() {
				Example(true)
			})
			Attribute("user_ids", ArrayOf(Integer), "ルームに入れるUserID", func() {
				Example([]int{1, 2})
			})
			Attribute("img", String, "画像データ", func() {
				Example("img.com")
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

	Action("exists", func() {
		Routing(GET("rooms/exists"))
		Description("DMの存在を確認する")
		Params(func() {
			Param("user_id", Integer, "ユーザーID")
			Required("user_id")
		})
		Response(OK, room)
		Response(NotFound)
		Response(InternalServerError)
	})

	Action("show", func() {
		Routing(GET("rooms/:id"))
		Description("ルームの詳細を取得する")
		Params(func() {
			Param("id", Integer, "ID")
		})
		Response(OK, roomUser)
		Response(NotFound)
		Response(BadRequest)
		Response(InternalServerError)
	})
})

var room = MediaType("application/vnd.room", func() {
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
	Attribute("img", String, "画像")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("is_group")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("img")
	})
	Required("id", "name", "is_group", "created_at", "updated_at")
})

var roomUser = MediaType("application/vnd.room_user", func() {
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
	Attribute("img", String, "画像")
	Attribute("users", CollectionOf(show_user), "ルームいるユーザー")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("is_group")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("img")
		Attribute("users")
	})
	Required("id", "name", "is_group", "created_at", "updated_at", "users")
})

var indexRoom = MediaType("application/vnd.index_room", func() {
	Description("ルームの表示")
	Attribute("room", room, "room")
	Attribute("is_open", Boolean, "開いたどうか")
	Attribute("last_text", String, "最後の内容")
	Attribute("count_user", Integer, "ルームに入っているユーザー数")
	View("default", func() {
		Attribute("room")
		Attribute("is_open")
		Attribute("last_text")
		Attribute("count_user")
	})
	Required("room", "is_open", "count_user")
})

var allRoomUser = MediaType("application/vnd.all_room_user", func() {
	Description("全てのルーム")
	Attribute("index_room", CollectionOf(indexRoom), "index_rooms")
	Attribute("next_id", Integer, "次取得するRoomID")
	View("default", func() {
		Attribute("index_room")
		Attribute("next_id")
	})
	Required("index_room")
})
