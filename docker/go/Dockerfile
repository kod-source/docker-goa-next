FROM golang:1.19
ENV GO111MODULE on
ENV ROOT=/go/src/app

WORKDIR ${ROOT}
# COPY go.mod go.sum ./

# 併せて air もインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air", "-c", ".air.toml"]
