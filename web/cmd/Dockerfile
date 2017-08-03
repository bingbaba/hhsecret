FROM golang:1.8.3-onbuild
MAINTAINER Xue Bing <xuebing1110@gmail.com>

RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# move to GOPATH
RUN mkdir -p /go/src/github.com/bingbaba/hhsecret
RUN cp -r /go/src/app/* $GOPATH/src/github.com/bingbaba/hhsecret

# WORKDIR
WORKDIR $GOPATH/src/github.com/bingbaba/hhsecret

RUN mkdir /etc/hhsecret
RUN cp web/cmd/conf.yml /etc/hhsecret/conf.yml
RUN go build -o hhsecret web/cmd/main.go

EXPOSE 8003
CMD ["hhsecret", "-c", "/etc/hhsecret/conf.yml"]