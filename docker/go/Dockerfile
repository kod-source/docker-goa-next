FROM golang:1.20
ENV GO111MODULE on
ENV ROOT=/go/src/app

WORKDIR ${ROOT}
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

# 併せて air もインストール
RUN go install github.com/cosmtrek/air@latest

COPY ./api ./

# airコマンドでGoファイルを起動
CMD ["air", "-c", ".air.toml"]
