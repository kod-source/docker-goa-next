package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var ctx context.Context
var testDB *sql.DB
var jst *time.Location

func TestMain(m *testing.M) {
	var cleanup func()
	ctx, testDB, cleanup = newTest()
	defer cleanup()

	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Exit(1)
	}

	m.Run()
}

func newTest() (context.Context, *sql.DB, func()) {
	if err := godotenv.Load("../../../.env"); err != nil {
		// CIで.envファイルはないので、ログを吐くのみにする
		log.Println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	dbName := "test_" + os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = user
	config.Passwd = password
	config.Addr = net.JoinHostPort(host, port)
	config.ParseTime = true

	// テスト用の新規データーベース作成
	func() {
		config := config.Clone()
		config.MultiStatements = true
		db0, err := sql.Open("mysql", config.FormatDSN())
		if err != nil {
			panic(err)
		}
		defer db0.Close()

		if _, err := db0.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE `%s`", dbName)); err != nil {
			panic(err)
		}
	}()

	// テーブル初期化
	func() {
		config := config.Clone()
		config.DBName = dbName
		config.MultiStatements = true
		db0, err := sql.Open("mysql", config.FormatDSN())
		if err != nil {
			panic(err)
		}
		defer db0.Close()

		// app/schema/tables.sql を参照する
		f, err := os.Open("../schema/schema.sql")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		buf, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		if _, err := db0.ExecContext(ctx, string(buf)); err != nil {
			panic(err)
		}
	}()

	config.DBName = dbName
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}

	return ctx, db, func() {
		// clean up
		_, _ = db.ExecContext(ctx, fmt.Sprintf("DROP DATABASE `%s`", dbName))
		_ = db.Close()
		cancel()
	}
}
