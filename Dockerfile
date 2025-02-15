# 使用官方的 Go 镜像作为构建环境
FROM golang:1.20 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o GoToDoList cmd/app/main.go

# 使用官方的 Alpine Linux 镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/GoToDoList .

# 复制配置文件
COPY configs /app/configs

# 暴露端口
EXPOSE 8080

# 运行应用程序
CMD ["./GoToDoList"]
