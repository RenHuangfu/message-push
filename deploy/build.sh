# 构建 manager 服务
docker build \
  -f app/manager/Dockerfile \
  -t manager-service:latest \
  .

# 构建 pusher 服务
docker build \
  -f app/pusher/Dockerfile \
  -t pusher-service:latest \
  .

# 构建 receiver 服务
docker build \
  -f app/receiver/Dockerfile \
  -t receiver-service:latest \
  .