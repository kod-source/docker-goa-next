package main

import (
	"database/sql"
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tanimutomo/sqlfile"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("DSN")+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filePath := flag.String("file_path", "", "file path")
	flag.Parse()
	if *filePath == "" {
		panic("引数を設定してください")
	}

	s := sqlfile.New()
	if err = s.File(*filePath); err != nil {
		panic(err)
	}
	_, err = s.Exec(db)
	if err != nil {
		panic(err)
	}
}
