FROM registry.cn-hangzhou.aliyuncs.com/5eplay/golang-alpine-1.19.0:v1.0.0 AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
#COPY service/bar-single/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/dio service/foo-multi/ctlcmd/ctl.go

FROM registry.cn-hangzhou.aliyuncs.com/5eplay/alpine:3.16.2
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/dio /app/dio
#COPY --from=builder /app/etc /app/etc
#COPY --from=builder /app/manifests /app/manifests

CMD ["./dio","ctl"]