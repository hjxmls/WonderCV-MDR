FROM golang:1.14.4

RUN apt-get update && apt-get upgrade -y && apt-get install -y vim

ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn
ENV GIN_MODE=release
ENV IM_DEBUG_MODE=false
RUN mkdir /root/app
WORKDIR $GOPATH/src/im-backend-gin
COPY . $GOPATH/src/im-backend-gin
RUN  go build -o entry.go ./main/main.go


EXPOSE 5959
ENTRYPOINT ["./entry.go"]