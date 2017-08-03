FROM golang:1.8.3-onbuild
MAINTAINER Xue Bing <xuebing1110@gmail.com>

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN mkdir /etc/hhsecret
RUN cp web/cmd/conf.yml /etc/hhsecret/conf.yml
RUN go build -o hhsecret web/cmd/main.go


EXPOSE 8003
CMD ["hhsecret", "-c", "/etc/hhsecret/conf.yml"]