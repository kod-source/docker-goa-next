package schema

import (
	"database/sql"
	"time"

	"github.com/shogo82148/myddlmaker"
)

// Thread スレッド
type Thread struct {
	// ID ...
	ID uint64 `ddl:",auto"`
	// UserID ...
	UserID uint64
	// RoomID ...
	RoomID uint64
	// Text ...
	Text string
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
	// Img 画像データ
	Img sql.NullString `ddl:",null,type=LONGTEXT"`
}

func (*Thread) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Thread) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_user_id", "user_id"),
		myddlmaker.NewIndex("idx_room_id", "room_id"),
	}
}

func (*Thread) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_thread_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"thread_room_id_constraint",
			[]string{"room_id"},
			"room",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}

func (*Thread) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		myddlmaker.NewUniqueIndex("user_room_id_index", "user_id", "room_id"),
	}
}
