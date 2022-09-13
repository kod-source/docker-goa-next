package schema

import (
	"database/sql"
	"os"
)

// NewDB mysqlのデータベースを起動する
func NewDB() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DSN")+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, nil
}
