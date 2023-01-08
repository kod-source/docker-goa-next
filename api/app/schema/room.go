package schema

import (
	"database/sql"
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
	// Img 画像データ
	Img sql.NullString `ddl:",null,type=LONGTEXT"`
}

func (*Room) PrimaryKey() *myddlmaker.PrimaryKey {
	return myddlmaker.NewPrimaryKey("id")
}
