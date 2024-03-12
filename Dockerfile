FROM 172.16.100.99:9006/golang:alpine AS builder

LABEL stage=gobuilder
ENV TZ Asia/Shanghai
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY ./build/drone-k8s /app/drone-k8s

# run step
FROM xxx.xxx.cn:9043/public/self-alpine

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#RUN apk update --no-cache
#RUN apk add --no-cache ca-certificates
#RUN apk add --no-cache tzdata
#COPY ./config /app/configabcd

ENV MODE_ENV test
WORKDIR /app
# copy bin from build step
COPY --from=builder /app/drone-k8s .

ENTRYPOINT ["/app/drone-k8s"]