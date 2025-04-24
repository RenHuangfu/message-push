package test

import (
	"bufio"
	"fmt"
	"net"
)

func Test(msg string) {
	conn, err := net.Dial("tcp", "10.63.5.206:8850")
	if err != nil {
		fmt.Println("连接失败:", err)
	}
	defer conn.Close()

	// 读取用户输入并发送
	writer := bufio.NewWriter(conn)

	_, err = writer.WriteString(msg + "\n")
	if err != nil {
		return
	}
	writer.Flush()
}
