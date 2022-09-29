package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shogo82148/schemalex-deploy/deploy"
	"github.com/shogo82148/schemalex-deploy/diff"
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
	s, err := db.LoadSchema(ctx)
	if err != nil {
		panic(err)
	}
	if err := diff.Strings(os.Stdout, s, newDDL, diff.WithTransaction(true)); err != nil {
		fmt.Println("エラーでうしょー")
		fmt.Println(err)
	}
	// plan, err := db.Plan(ctx, s)
	// if err != nil {
	// 	fmt.Println("エラー")
	// 	panic(err)
	// }
	// fmt.Println(plan)
}

func readFile(filePath string) (string, error) {
	// ファイルをOpenする
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	return string(b), nil
}
