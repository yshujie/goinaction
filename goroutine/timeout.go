package goroutine_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/**
✅ 题目：A4 超时控制
在一个消费者等待生产者的场景中，增加超时控制。如果在指定时间内没有收到数据，就打印超时消息并退出。

函数格式：
func timeout(timeout time.Duration)

输入：
timeout time.Duration —— 最大等待时间

输出：
收到数据时打印数据；超时时打印超时提示。

考点：
- select + time.After 实现超时机制
- channel 接收超时模式
- goroutine 与超时退出的配合
*/

func Timeout(timeout time.Duration) {
	// 1. 启动 goroutine，等待输入数据
	// 输入 goroutine：读到一行就投递一次结果（缓冲=1，确保不阻塞）
	inputChannel := make(chan string, 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		// 去掉末尾的换行符，写入 channel
		inputChannel <- strings.TrimSpace(inputStr)
	}()

	// 2. 启动倒计时
	remaining := int(timeout / time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 3. 启动延迟触发器，倒计时 timeout 秒
	downline := time.After(timeout)

	// 4. select 多路选择
	for {
		select {
		case <-ticker.C: // 倒计时
			if remaining >= 0 {
				fmt.Printf("\r请输入内容，按回车键结束(倒计时：%vs)：", remaining)
				remaining--
			}
		case <-downline:
			fmt.Printf("\n在 %v 内没有输入，自动退出\n", timeout)
			return
		case value := <-inputChannel:
			fmt.Printf("输入信息：%v\n", value)
			return
		}
	}
}
