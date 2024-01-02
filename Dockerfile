FROM golang:1.19 AS builder

# Set the Current Working Directory inside the container
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 移动到工作目录 /build
WORKDIR /build

# 将代码复制到容器中
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o acs .


# 运行二进制文件
FROM alpine:latest

# 安装 Nginx
RUN apk update && apk add nginx

# 设置时区
ENV TZ=Asia/Shanghai

# 安装 tzdata 包
RUN apk add tzdata \
    && echo "${TZ}" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && rm /var/cache/apk/*

# 移动到 Nginx 默认的静态文件目录
WORKDIR /usr/share/nginx/html

COPY --from=builder /build/web/acs_admin/dist .
COPY --from=builder /build/web/front_nginx.conf /etc/nginx/http.d/default.conf


# 移动到根目录
WORKDIR /run

# 将文件拷贝到当前目录
COPY --from=builder /build/acs ./acs
COPY --from=builder /build/.env ./.env

# 声明需要暴露的端口 实际需要根据配置文件修改
# http服务端口
EXPOSE 9528
# gRPC 服务端口
EXPOSE 9529
# 后台管理端口
EXPOSE 80

# 执行权限
RUN chmod -R 755 ./acs

# 需要启动的命令
ENTRYPOINT ["sh", "-c", "nginx && ./acs"]


