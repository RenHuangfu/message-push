docker build -t static-server:latest .
docker run -d -p 7788:7788 --name my-server static-server