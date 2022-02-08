FROM golang:1.17 as builder

ENV GOSUMDB=off GOPROXY=https://goproxy.cn GO111MODULE=on

# 增加分层，减少下载依赖包
WORKDIR /go/cache

# go 依赖
COPY ["go.mod","go.sum","./"]

RUN go mod download  && mkdir -p /go/src/project


WORKDIR /go/src/project

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update  \
    && apk add --no-cache ca-certificates  \
    && apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

EXPOSE 8888

WORKDIR /root


COPY --from=builder /go/src/project/main .






ENTRYPOINT ["/root/main"]
