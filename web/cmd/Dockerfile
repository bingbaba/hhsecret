FROM golang:1.8.3-onbuild
MAINTAINER Xue Bing <xuebing1110@gmail.com>

#EXPOSE 1318

RUN cd web/cmd
RUN go build -o hhsecret ./main.go

CMD ["hhsecret", "-c", "./conf.yml"]