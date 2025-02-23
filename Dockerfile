# 第一步编译go文件
FROM golang:1.24-alpine AS builder

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目的所有文件
COPY . .

# 构建可执行文件
RUN go build -o main ./src/main.go

# 第二步生成镜像
FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/main /app/main
COPY ./src/app.ini /app/app.ini

# 安装必要的运行时依赖
RUN apk add --no-cache ca-certificates

# 暴露应用程序端口（根据你的应用程序需要调整端口）
EXPOSE 8000

# 运行可执行文件
CMD ["./main"]
