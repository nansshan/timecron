# 使用官方 Golang 镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目源代码
COPY . .

# 整理依赖，确保 go.mod 和 go.sum 是最新的
RUN go mod tidy

# 编译项目
# CGO_ENABLED=0 禁用 CGO，以便在 Alpine 这种不含 C 库的环境中运行
# GOOS=linux 指定编译为 Linux 可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o timecron main.go

# 使用更小的基础镜像构建最终镜像
FROM alpine:latest

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建环境中复制编译后的二进制文件
COPY --from=builder /app/timecron .

# 复制配置文件
COPY config.json .

# 复制静态文件
COPY static ./static


# 暴露端口 (根据 api.http, 项目运行在 3005 端口)
EXPOSE 3005

# 运行命令
CMD ["./timecron"] 