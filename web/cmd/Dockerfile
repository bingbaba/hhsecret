FROM golang:1.8.3-onbuild
MAINTAINER Xue Bing <xuebing1110@gmail.com>

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# move to GOPATH
RUN mkdir -p /app
RUN mkdir -p /go/src/github.com/bingbaba/hhsecret
COPY . $GOPATH/src/github.com/bingbaba/hhsecret/
COPY web/html /app/html/
WORKDIR $GOPATH/src/github.com/bingbaba/hhsecret

# example config
# RUN mkdir /etc/hhsecret
# RUN cp web/cmd/conf.yml /app/conf.yml

# build
RUN go build -o /app/hhsecret web/cmd/main.go

EXPOSE 8003
WORKDIR /app
CMD ["/app/hhsecret", "-c", "/etc/hhsecret/conf.yml"]