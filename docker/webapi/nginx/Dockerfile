# ベースイメージ
FROM nginx:mainline-alpine

# Nginxの設定ファイルをコピー
COPY nginx.conf /etc/nginx/nginx.conf
COPY app.conf /etc/nginx/conf.d/app.conf

# コンテナ起動時にNginxを実行
CMD ["nginx", "-g", "daemon off;"]

