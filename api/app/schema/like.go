package schema

import (
	"github.com/shogo82148/myddlmaker"
)

//go:generate go run -tags myddlmaker gen/main.go

type Like struct {
	ID     uint64 `ddl:",auto"`
	PostID uint64
	UserID uint64
}

func (*Like) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Like) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"u_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade).OnDelete(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"p_id_constraint",
			[]string{"post_id"},
			"post",
			[]string{"id"},
		).OnUpdate(myddlmaker.ForeignKeyOptionCascade).OnDelete(myddlmaker.ForeignKeyOptionCascade),
	}
}

func (*Like) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		// UNIQUE INDEX `idx_name` (`name`)
		myddlmaker.NewUniqueIndex("post_user_id_index", "post_id", "user_id"),
	}
}
