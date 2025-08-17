package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：A3 多消费者模型
实现一个生产者 + 多个消费者的模型，生产者通过有缓冲 channel 发送一批任务（例如整数），多个消费者并发接收并处理任务。

函数格式：
func multiConsumer(workerCount int, tasks []int)

输入：
workerCount int —— 消费者数量
tasks []int —— 任务数据列表

输出：
消费者编号 + 处理的任务值

考点：
- 有缓冲 channel 的使用
- 多 goroutine 消费同一 channel
- WaitGroup 等待所有消费者完成
- 任务分发的竞争特性
*/

func MultiConsumer(workerCount int, tasks []int) {
	// 非法的 worker 数量
	if workerCount <= 0 {
		return
	}

	// 创建有缓冲的通道
	ch := make(chan int, 3)

	// 创建 wait group
	var wg sync.WaitGroup
	wg.Add(workerCount)

	// 创建消费者
	for i := 1; i <= workerCount; i++ {
		go func(i int) {
			defer wg.Done()

			// 使用 for range 在 channel 中获取数据，直至 channel 中取空了为止
			for item := range ch {
				fmt.Printf("in [%v] worker, get %v param \n", i, item)
			}
		}(i)
	}

	// 生产者：向 channel 中写数据
	for _, item := range tasks {
		ch <- item // 队列满时会对生产者施加背压
	}
	// 发送完后关闭通道，广播“没有更多任务”
	close(ch)

	wg.Wait()
}
