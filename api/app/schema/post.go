package schema

import (
	"time"

	"github.com/shogo82148/myddlmaker"
)

//go:generate go run -tags myddlmaker gen/main.go

type Post struct {
	ID        uint64 `ddl:",auto`
	UserID    uint64
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Img       []byte `ddl:",null"`
}

func (*Post) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Post) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade).OnDelete(myddlmaker.ForeignKeyOptionCascade),
	}
}
