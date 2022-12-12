package schema

import (
	"github.com/shogo82148/myddlmaker"
)

type Like struct {
	ID     uint64 `ddl:",auto"`
	PostID uint64
	UserID uint64
}

func (*Like) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*Like) Indexes() []*myddlmaker.Index {
	return []*myddlmaker.Index{
		myddlmaker.NewIndex("idx_user_id", "user_id"),
		myddlmaker.NewIndex("idx_post_id", "post_id"),
	}
}

func (*Like) ForeignKeys() []*myddlmaker.ForeignKey {
	return []*myddlmaker.ForeignKey{
		myddlmaker.NewForeignKey(
			"u_id_constraint",
			[]string{"user_id"},
			"user",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),

		myddlmaker.NewForeignKey(
			"p_id_constraint",
			[]string{"post_id"},
			"post",
			[]string{"id"},
		).OnDelete(myddlmaker.ForeignKeyOptionCascade).OnUpdate(myddlmaker.ForeignKeyOptionCascade),
	}
}

func (*Like) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		// UNIQUE INDEX `idx_name` (`name`)
		myddlmaker.NewUniqueIndex("post_user_id_index", "post_id", "user_id"),
	}
}
