package schema

import (
	"time"

	"github.com/shogo82148/myddlmaker"
)

type Post struct {
	ID        uint64 `ddl:",auto"`
	UserID    uint64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Img       string `ddl:",null,type=LONGTEXT"`
}

func (*Post) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Post) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_user_id", "user_id"),
	}
}

func (*Post) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}
