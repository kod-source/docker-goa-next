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

	"github.com/caarlos0/env"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kod-source/docker-goa-next/app/datastore/testdata"
)

var ctx context.Context
var testDB *sql.DB
var jst *time.Location
var now time.Time

type dbConfig struct {
	DatabaseName     string `env:"MYSQL_DATABASE,required"`
	DatabaseUser     string `env:"MYSQL_ROOT_USER,required"`
	DatabasePassword string `env:"MYSQL_ROOT_PASSWORD,required"`
	DatabasePort     string `env:"MYSQL_LOCAL_PORT" envDefault:"3306"`
	DatabaseHost     string `env:"MYSQL_LOCAL_HOST,required"`
}

func getDBConfig() (*dbConfig, error) {
	cfg := dbConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func TestMain(m *testing.M) {
	var cleanup func()
	ctx, testDB, cleanup = newTest()
	defer cleanup()

	if err := testdata.UserSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.PostSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.CommentSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.LikeSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.RoomSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.UserRoomSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.ThreadSeed(ctx, testDB); err != nil {
		panic(err)
	}
	if err := testdata.ContentSeed(ctx, testDB); err != nil {
		panic(err)
	}
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		os.Exit(1)
	}
	now = time.Date(2022, 1, 1, 0, 0, 0, 0, jst)

	m.Run()
}

func newTest() (context.Context, *sql.DB, func()) {
	if err := godotenv.Load("../../../.env"); err != nil {
		// CIで.envファイルはないので、ログを吐くのみにする
		log.Println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	dbConf, err := getDBConfig()
	if err != nil {
		panic(err)
	}
	dbName := "test_" + dbConf.DatabaseName

	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = dbConf.DatabaseUser
	config.Passwd = dbConf.DatabasePassword
	config.Addr = net.JoinHostPort(dbConf.DatabaseHost, dbConf.DatabasePort)
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
