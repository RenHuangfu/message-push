# 第一阶段：构建Go二进制
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制源代码
COPY main.go .

# 编译静态二进制文件（禁用CGO）
RUN CGO_ENABLED=0 go build -o static-server main.go

# 第二阶段：创建最小化运行镜像
FROM alpine:latest

WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/static-server .

# 创建默认静态文件目录
RUN mkdir -p /app/static

# 设置默认的容器入口点和参数
ENTRYPOINT ["./static-server"]
CMD ["-dir", "/app/static", "-port", "7788"]

# 暴露默认端口
EXPOSE 7788