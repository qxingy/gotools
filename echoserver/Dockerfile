FROM golang:alpine AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.io,https://goproxy.cn,direct \
        && go generate ./... \
        && go build -v -o app .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/app .

EXPOSE 8080

CMD ["./app"]