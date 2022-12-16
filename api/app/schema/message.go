package schema

import (
	"database/sql"
	"time"

	"github.com/shogo82148/myddlmaker"
)

// Message スレッドの返信
type Message struct {
	// ID ...
	ID uint64 `ddl:",auto"`
	// UserID ...
	UserID uint64
	// ThreadID ...
	ThreadID uint64
	// Text ...
	Text string
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
	// Img 画像データ
	Img sql.NullString `ddl:",null,type=LONGTEXT"`
}

func (*Message) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Message) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_user_id", "user_id"),
		myddlmaker.NewIndex("idx_thread_id", "thread_id"),
	}
}

func (*Message) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_message_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"thread_message_id_constraint",
			[]string{"thread_id"},
			"thread",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}

func (*Message) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		myddlmaker.NewUniqueIndex("user_thread_id_index", "user_id", "thread_id"),
	}
}
