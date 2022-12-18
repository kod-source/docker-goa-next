package schema

import (
	"time"

	"github.com/shogo82148/myddlmaker"
)

type Room struct {
	// ID ...
	ID uint64 `ddl:",auto"`
	// Name 部屋名
	Name string
	// IsGroup グループかDMかの判定
	IsGroup bool
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
}

func (*Room) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}
