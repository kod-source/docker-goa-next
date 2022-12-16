package schema

import (
	"database/sql"
	"time"

	"github.com/shogo82148/myddlmaker"
)

// UserRoom チャットの参加者
type UserRoom struct {
	// ID ...
	ID uint64 `ddl:",auto"`
	// UserID ...
	UserID uint64
	// RoomID ...
	RoomID uint64
	// LastReadAt このルームの最後に既読をつけた日
	LastReadAt sql.NullTime `ddl:",null"`
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
}

func (*UserRoom) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*UserRoom) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_user_id", "user_id"),
		myddlmaker.NewIndex("idx_room_id", "room_id"),
	}
}

func (*UserRoom) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_room_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"room_id_constraint",
			[]string{"room_id"},
			"room",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}

func (*UserRoom) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		myddlmaker.NewUniqueIndex("user_room_id_index", "user_id", "room_id"),
	}
}
