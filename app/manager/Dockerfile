# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
COPY third_party ./third_party

ENV GOPROXY=https://goproxy.cn,direct
# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 构建二进制
RUN CGO_ENABLED=0 GOOS=linux go build -o /manager-service ./app/manager/service/cmd/server

# 运行阶段
FROM alpine:latest

# 安装必要的运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /root/

# 复制配置文件
COPY --from=builder /app/app/manager/service/configs/config.yaml .

# 复制可执行文件
COPY --from=builder /manager-service .

# 启动命令
CMD ["./manager-service", "-conf", "config.yaml"]