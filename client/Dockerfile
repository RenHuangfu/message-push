# 修改后的 Dockerfile
# 第一阶段：构建Go二进制
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制所有源代码（包括静态文件）
COPY . .

# 编译静态二进制文件（禁用CGO）
RUN CGO_ENABLED=0 go build -o static-server main.go

# 第二阶段：创建最小化运行镜像
FROM alpine:latest

WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/static-server .
# 关键修改：复制静态文件到容器内的静态目录
COPY --from=builder /app/ws.html /app/ws.js /app/static/

# 暴露默认端口
EXPOSE 7788

ENTRYPOINT ["./static-server"]
CMD ["-dir", "/app/static", "-port", "7788"]