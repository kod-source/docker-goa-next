# apiサーバーとクラインを同時に起動する
local:
	docker compose up

# designファイルの実行 go generate ./...
design:
	docker compose run api go get github.com/shogo82148/goa-v1/...@v1 && docker compose run api go generate ./... && docker-compose run api go mod tidy
