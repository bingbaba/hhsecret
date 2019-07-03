FROM golang:1.12.6-alpine3.10
MAINTAINER Xue Bing <xuebing1110@gmail.com>

# repo
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
RUN echo "http://mirrors.aliyun.com/alpine/v3.10/main/" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.10/community/" >> /etc/apk/repositories

# timezone
RUN apk update
RUN apk add --no-cache tzdata \
    && echo "Asia/Shanghai" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk add curl

# move to GOPATH
RUN mkdir -p /app
#RUN mkdir -p /go/src/github.com/bingbaba/hhsecret
#COPY . $GOPATH/src/github.com/bingbaba/hhsecret/
COPY web/html /app/html/
#WORKDIR $GOPATH/src/github.com/bingbaba/hhsecret

# example config
# RUN mkdir /etc/hhsecret
# RUN cp web/cmd/conf.yml /app/conf.yml

# build
ENV GO111MODULE=on
RUN go mod download
RUN go build -o /app/hhsecret web/cmd/main.go

EXPOSE 8003
WORKDIR /app
CMD ["/app/hhsecret", "-c", "/etc/hhsecret/conf.yml"]
