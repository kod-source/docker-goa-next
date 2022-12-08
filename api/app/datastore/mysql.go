package datastore

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// NewDB mysqlのデータベースを起動する
func NewDB(conf *mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
