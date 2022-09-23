FROM golang:1.18 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY config/ config/
#COPY go.sum go.sum
RUN go mod download

COPY main.go main.go

ARG VERSION=undefined

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build \
    -ldflags "-X main.Version=$VERSION" \
    -a \
    -o bin/go-imdb main.go

FROM debian:bookworm-slim

WORKDIR /

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN addgroup --gid 901 go-imdb && adduser --uid 901 --gid 901 --home /home/go-imdb go-imdb

COPY --from=builder /workspace/bin/go-imdb .
USER go-imdb

ENTRYPOINT ["/go-imdb"]
