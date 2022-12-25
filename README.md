# docker-goa-next

## 起動方法
```
docker compose build
```
yarnで、package.jsonに基づいて、依存パッケージを管理をしていて、node_modules内のファイルはGit管理していないため、初回起動時のみ実行が必要
```
docker compose run client yarn install
```
```
docker compose up -d
```

## マイグレーション方法
使用ツール
- git-schemalex

### インストール方法
```
go install github.com/schemalex/git-schemalex/cmd/git-schemalex
```
### マイグレーション実行方法
```
git schemalex -schema api/app/schema/schema.sql -dsn "$user:$password@tcp($host:$port)/$db_name" -deploy
```
* 値は`.env`を参照すること
