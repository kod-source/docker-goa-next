package schema

import (
	"time"

	"github.com/shogo82148/myddlmaker"
)

type Comment struct {
	ID        uint64 `ddl:",auto"`
	PostID    uint64
	UserID    uint64
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Img       string `ddl:",null,type=LONGTEXT"`
}

func (*Comment) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Comment) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_post_id", "post_id"),
		myddlmaker.NewIndex("idx_user_id", "user_id"),
	}
}

func (*Comment) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"user_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"post_id_constraint",
			[]string{"post_id"},
			"post",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}
