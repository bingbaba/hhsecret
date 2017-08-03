FROM golang:1.8.3-onbuild
MAINTAINER Xue Bing <xuebing1110@gmail.com>

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# move to GOPATH
RUN mkdir -p /go/src/github.com/bingbaba/hhsecret
RUN cp -r /go/src/app/* $GOPATH/src/github.com/bingbaba/hhsecret

# RUN mkdir /etc/hhsecret
# RUN cp web/cmd/conf.yml /app/conf.yml
RUN go build -o /app/hhsecret web/cmd/main.go

# WORKDIR
RUN mkdir -p /app
RUN cp -r /go/src/app/web/html /app/
WORKDIR /app

EXPOSE 8003
CMD ["/app/hhsecret", "-c", "/app/conf.yml"]