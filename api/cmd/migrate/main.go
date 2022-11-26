package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/schemalex/schemalex/diff"
	"github.com/shogo82148/schemalex-deploy/deploy"
)

func main() {
	ctx := context.Background()
	newDDL, err := readFile("app/schema/schema.sql")
	if err != nil {
		panic(err)
	}

	db, err := deploy.Open("mysql", "user:password@tcp(localhost:3306)/db_goa_next?parseTime=true")
	// db, err := deploy.Open("mysql", "user:password@tcp(db-goa:3306)/db_goa_next")
	// `schemalex_revision`テーブルの更新処理
	// if err := db.Import(ctx, newDDL); err != nil {
	// 	panic(err)
	// }

	// 現状のCREATE TABLEを取得する
	s, err := db.LoadSchema(ctx)
	if err != nil {
		panic(err)
	}

	// sqlの差分を表示する
	diff.Strings(os.Stdout, s, newDDL, diff.WithTransaction(true))
	// if err := diff.Strings(os.Stdout, s, newDDL, diff.WithTransaction(true)); err != nil {
	// 	fmt.Println("エラーでうしょー")
	// 	fmt.Println(err)
	// }

	plan, err := db.Plan(ctx, s)
	if err != nil {
		fmt.Println("エラー")
		panic(err)
	}
	fmt.Println(plan)
}

func readFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	return string(b), nil
}
