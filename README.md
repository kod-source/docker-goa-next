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
