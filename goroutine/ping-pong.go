package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：A1 Ping-Pong（双 goroutine 交替通信）
实现两个 goroutine 通过 channel 往返传递一个整数，每次传递时将该整数加 1，并打印当前 goroutine 的名称和数值。
数值达到指定上限时，停止通信并退出程序。

函数格式：
func pingPong(limit int)

输入：5
输出：A1B2A3B4A5

考点：
- goroutine 的创建与启动
- 无缓冲 channel 的同步机制
- 多 goroutine 间的消息传递
- 程序退出条件的控制
*/

/*
*
思考：
- 要求：两个 goroutine、一个 channel
- goroutine 设计：
  - goroutine A：作为 channel 的消费者，输出 A<i>，进度++，并将控制权转交给 goroutine B；判断进度，到达上限时结束任务，关闭通道
  - goroutine B：作为 channel 的消费者，输出 B<i>，进度++，并将控制全转交给 goroutine A；判断进度，到达上限时结束任务，关闭通道

- channel 设计：
  - 打印控制器，包含 打印者、进度 字段

- 谁生产消息
  - 主 goroutine 生产第一条消息
  - A goroutine 生产 “让 B goroutine 工作” 的消息
  - B goroutine 生产 “让 A goroutine 工作” 的消息

- 谁消费消息
  - A goroutine 消费 “A goroutine 工作” 的消息
  - B goroutine 消费 “B goroutine 工作” 的消息

- 谁关闭
  - A goroutine 判断到达上限，就关闭
  - B goroutine 判断到达上限，就关闭
*/
func PingPong(limit int) {
	if limit <= 0 {
		return
	}

	// 输出器
	type printer int
	const (
		aPrinter = +iota
		bPrinter
	)

	// 打印控制器
	type ctr struct {
		p printer
		i int
	}

	// 创建无缓冲的通道，作为共享内存，记录打印控制器
	channel := make(chan ctr)

	// 设置主进程等待
	var wait sync.WaitGroup
	// 因有 A、B 两个 goroutine，固设置等待数量为 2
	wait.Add(2)

	// goroutine A
	go func() {
		defer wait.Done()

		// 通道关闭自动退出循环
		for c := range channel {
			// 若打印器不是 A，则将消息塞回去
			if c.p != aPrinter {
				channel <- c
				continue
			}

			// 打印器为 A， 则打印
			fmt.Printf("A%d", c.i)

			// 若到达上限，则关闭通道、退出
			if c.i >= limit {
				close(channel)
				return
			}

			// 将控制器交给 goroutine B
			c.i++
			c.p = bPrinter

			channel <- c
		}
	}()

	// goroutine B
	go func() {
		defer wait.Done()

		// 通道关闭自动退出循环
		for c := range channel {
			// 若打印器不是 B，则将消息塞回去
			if c.p != bPrinter {
				channel <- c
				continue
			}

			// 打印器为 B， 则打印
			fmt.Printf("B%d", c.i)

			// 若到达上限，则关闭通道、退出
			if c.i >= limit {
				close(channel)
				return
			}

			// 将控制器交给 goroutine A
			c.i++
			c.p = aPrinter

			channel <- c
		}
	}()

	// 主协程启动打印，并等待子协程
	channel <- ctr{p: aPrinter, i: 1}
	wait.Wait()
}
