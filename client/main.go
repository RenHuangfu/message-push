package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "7788", "服务器监听端口")
	dir := flag.String("dir", "./", "静态文件目录")
	flag.Parse()

	// 静态文件服务
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	// 启动服务器
	log.Printf("服务器启动，监听端口 %s，静态文件目录 %s", *port, *dir)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
