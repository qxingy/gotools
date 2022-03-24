FROM golang:1.16

COPY . /go/src

WORKDIR /go/src

RUN  GOPROXY=https://goproxy.io,https://goproxy.cn,direct \
         && go build -v -o /go/bin/. . 

WORKDIR /go/bin

EXPOSE 8080

CMD ["./echoserver"]