package goroutine_test

import (
	"fmt"
	"sync"
	"time"
)

/**
✅ 题目：A6 关闭 channel 的信号广播
实现一个生产者向多个消费者广播完成信号的程序。生产者完成任务后关闭 channel，所有消费者检测到 channel 被关闭后自动退出。

函数格式：
func broadcastClose(workerCount int)

输入：
workerCount int —— 消费者数量

输出：
消费者接收到数据的日志；接收到关闭信号后的退出提示。

考点：
- for range 接收 channel 数据直到关闭
- close(channel) 的广播特性
- 多消费者安全退出
*/

/*
*
思考：
- 通过事件广播通知所有消费者关闭
*/
func BroadcastClose(workerCount int) {
	if workerCount <= 0 {
		return // 没有消费者就不启动，避免生产后没人读而阻塞
	}

	data := make(chan int, 8) // 缓冲不是必须；小缓冲有助于平滑吞吐

	var pwg sync.WaitGroup
	pwg.Add(1)
	// 生产者：负责发送完所有数据后 close(data)
	go func() {
		defer pwg.Done()
		defer close(data) // ✅ “谁生产谁关闭”——关闭即广播完结

		for i := 1; i <= 20; i++ {
			data <- i
			// 可选：模拟生产耗时
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// 启动多个消费者：读取到通道关闭为止
	var cwg sync.WaitGroup
	cwg.Add(workerCount)
	for id := 0; id < workerCount; id++ {
		go func(id int) {
			defer cwg.Done()
			for v := range data { // ✅ 通道关闭且读尽后自然退出
				fmt.Printf("[consumer %02d] recv: %d\n", id, v)
				// 可选：模拟消费耗时
				time.Sleep(50 * time.Millisecond)
			}
			fmt.Printf("[consumer %02d] channel closed, exit\n", id)
		}(id)
	}

	// 等生产者结束（可选），再等所有消费者读尽退出
	pwg.Wait()
	cwg.Wait()
}
