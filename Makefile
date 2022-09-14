# apiサーバーとクラインを同時に起動する
local:
	docker compose up

# designファイルの実行 go generate ./...
design:
	cd api && go get github.com/shogo82148/goa-v1/...@v1 && go generate ./... && go mod tidy && cd ../

# マイグレーションの実行
migrate:
	docker compose exec api go run cmd/sql.go -file_path=app/schema/ddl.sql

# 初期値の登録
seed:
	docker compose exec api go run cmd/sql.go -file_path=app/schema/seed.sql
