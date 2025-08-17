package goroutine_test

import (
	"fmt"
	"sync"
	"time"
)

/**
✅ 题目：A2 Channel 阻塞演示（生产者阻塞）
实现一个只有生产者 goroutine 往无缓冲 channel 中发送数据的程序，消费者延迟启动，用于观察和验证生产者在没有接收方时的阻塞行为。

函数格式：
func producerBlocking()

输入：
无

输出：
打印生产者开始发送、阻塞等待，以及消费者接收数据的时间点。

考点：
- 无缓冲 channel 的发送阻塞特性
- goroutine 执行时序
- 使用 time.Sleep 模拟启动延迟
*/

func ProducerBlocking() {
	// 创建无缓冲的通道
	ch := make(chan struct{})

	var wait sync.WaitGroup
	wait.Add(2)

	startTime := time.Now()
	fmt.Println("[master] master goroutine start, time: ", startTime)

	// 生产者 goroutine
	go func() {
		defer wait.Done()

		time0 := time.Now()
		fmt.Println("[producer] in producer goroutine , time : ", time0, "since: ", time.Since(startTime))

		timeBefor := time.Now()
		fmt.Println("[producer] before msg into channel , time : ", timeBefor)
		ch <- struct{}{}
		fmt.Println("[producer] after  msg into channel , since: ", time.Since(timeBefor))
	}()

	// 消费者 goroutine
	go func() {
		defer wait.Done()

		time1 := time.Now()
		fmt.Println("[consumer] in consumer goroutine, time: ", time1, "since: ", time.Since(startTime))

		timeBefor := time.Now()
		time.Sleep(5 * time.Second)
		fmt.Println("[consumer] before get msg from channel, time: ", timeBefor)
		<-ch
		fmt.Println("[consumer] after get msg from channel, time: ", time.Now(), "since: ", time.Since(timeBefor))
	}()

	wait.Wait()
}
