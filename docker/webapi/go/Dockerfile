FROM golang:1.20 as dev
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

FROM golang:1.20 as builder
ENV ROOT=/go/src/app

WORKDIR ${ROOT}
COPY ./api/go.mod ./api/go.sum ./

RUN go mod download

COPY ./api ${ROOT}

RUN go build -o /app ${ROOT}/cmd/webapi/main.go

FROM gcr.io/distroless/base:nonroot
COPY --chown=nonroot:nonroot --from=builder /app /
CMD ["/app"]
