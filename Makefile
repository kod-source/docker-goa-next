help: ## この文章を表示します。
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: local
local: ## apiサーバーとクライアントを同時に起動する
	docker compose up -d db api client

.PHONY: desing
design: ## designファイルの実行 go generate ./...
	cd api && go get github.com/shogo82148/goa-v1/...@v1 && go generate ./... && go mod tidy && cd ../

.PHONY: seed
seed: ## 初期値の登録
	docker compose exec api go run cmd/sql.go -file_path=app/schema/seed.sql

.PHONY: test_api
test_api: ## apiのテスト

.PHONY: schemaspy
schemaspy: ## schemaspyの更新
	docker compose up schemaspy

.PHONY: view_schemaspy
view_schemaspy: ## schemaspyをブラウザで確認する
	docker compose up -d nginx_schemaspy
