package schema

import (
	"database/sql"
	"fmt"
	"os"
)

// NewDB mysqlのデータベースを起動する
func NewDB() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DRIVER"), os.Getenv("DSN")+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	sqlFile, err := os.ReadFile("app/schema/ddl.sql")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}
