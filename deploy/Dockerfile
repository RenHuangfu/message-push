# 构建阶段
FROM golang:1.23-alpine AS builder

# 定义构建参数
ARG SERVICE_NAME

# 设置工作目录
WORKDIR /app

# 复制依赖文件
COPY go.mod go.sum ./
COPY third_party ./third_party

# 设置GOPROXY以加速依赖下载
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download

# 复制整个项目
COPY . .

# 构建二进制文件（通过SERVICE_NAME动态指定路径）
RUN CGO_ENABLED=0 GOOS=linux go build -o /${SERVICE_NAME}-service ./app/${SERVICE_NAME}/service/cmd/server

# 运行阶段
FROM alpine:latest

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /root/

# 再次声明构建参数（仅在构建阶段传递，运行阶段需要重新声明）
ARG SERVICE_NAME

# 复制配置文件
COPY --from=builder /app/app/${SERVICE_NAME}/service/configs/config.yaml .

# 复制可执行文件
COPY --from=builder /${SERVICE_NAME}-service .

# 启动服务
CMD ["./${SERVICE_NAME}-service", "-conf", "config.yaml"]