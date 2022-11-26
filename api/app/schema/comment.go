package schema

import (
	"time"

	"github.com/shogo82148/myddlmaker"
)

//go:generate go run -tags myddlmaker gen/main.go

type Comment struct {
	ID        uint64 `ddl:",auto"`
	PostID    uint64
	UserID    uint64
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Img       []byte `ddl:",null"`
}

func (*Comment) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Comment) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade).OnDelete(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"post_id_constraint",
			[]string{"post_id"},
			"post",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade).OnDelete(myddlmaker.ForeignKeyOptionCascade),
	}
}
