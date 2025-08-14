package goroutine_test

import (
	"fmt"
	"io"
	"net"
	"time"
)

// clock 是一个定时报告时间的 TCP 服务器

func Clock() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("listen error: %v \n", err)
	}

	// 无限循环
	for {
		// 阻塞，等待请求的到来
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listen err: %v \n", err)
			continue
		}

		// 启动新协程处理，支持并发的请求处理
		go handleConn(conn)
	}
}

// 处理请求
func handleConn(c net.Conn) {
	defer c.Close()

	// 循环 10 次
	i := 10
	for i > 0 {
		// 向请求写入当前时间
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("write time err, conn is %v \n", c)
		}

		time.Sleep(1 * time.Second)
		i--
	}
}
