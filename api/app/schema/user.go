package schema

import (
	"database/sql"
	"time"

	"github.com/shogo82148/myddlmaker"
)

//go:generate go run -tags myddlmaker gen/main.go

type User struct {
	ID        uint64 `ddl:",auto"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Avatar    sql.NullString `ddl:",null,type=LONGTEXT"`
}

func (*User) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}

func (*User) UniqueIndexes() []*myddlmaker.UniqueIndex {
	return []*myddlmaker.UniqueIndex{
		myddlmaker.NewUniqueIndex("idx_email", "email"),
	}
}
