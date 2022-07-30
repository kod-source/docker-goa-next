FROM golang:1.18
ENV GO111MODULE on
ENV ROOT=/go/src/app

WORKDIR ${ROOT}
# COPY go.mod go.sum ./

# 併せて air もインストール
RUN go install github.com/cosmtrek/air@v1.27.3 \
    && go install github.com/shogo82148/goa-v1/...@v1
# RUN go mod download \
#     && go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
# CMD ["./scripts/go-start.sh"]
CMD ["air", "-c", ".air.toml"]
