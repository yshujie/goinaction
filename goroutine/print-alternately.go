package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：交替打印数字和字母
用一个无缓冲通道实现两个 goroutine 的严格交替打印：


输入：N=5
输出：A1B1A2B2A3B3A4B4A5B5

注意：
- A 打印在先，B 打印在后
- 只能一个通道传令牌；严禁向“自己负责接收”的通道再发送。
- 用 WaitGroup 收尾；退出无悬挂、无死锁；go test -race 通过。
*/

/*
*
双通道版

思考：
- goroutine：goroutine 的本质是 woker，作用是接收消息、进行打印，所以当前需要两个 goroutine，分别打印 A<i> 和 B<i>
  - goroutine A：监听 channel A 消息，打印权限给 A，则去打印
  - goroutine B：监听 channel B 消息，打印权限给 B，则去打印

- channel：由于消费者不能自己生产消息（会造成死锁），所以要创建两个 channel，分别控制 A 和 B 两个 goroutine 的打印权限
  - channel A：goroutine A 的打印控制器，每次向 channel A 中发一条消息，代表让 goroutine A 打印一次
  - channel B：goroutine B 的打印控制器，每次向 channel B 中发一条消息，代表让 goroutine B 打印一次

-谁发送：
  - 在 主 goroutine 中生产：一个 for 循环，每次向 channel A、channel B 中生产消息
  - 不对，这样的话虽然能控制 1,2,3 ... 这样的顺序，但无法控制 A 和 B 谁先谁后
  - 主 goroutine 触发 A 的起始，A goroutine 运行一次后交接给 B goroutine，B goroutine 运行一次后交接给 A goroutine
  - 主 goroutine 最后控制 close

-谁接收：
  - channel A 的消费者：goroutine A
  - channel B 的消费者：goroutine B

-何时关闭：
  - 当 n > limit 时，由生产者关闭通道

-何时阻塞
  - 无缓冲的 channel，一次收发配对完成后双方解除阻塞，下一次操作才会再次阻塞等待对端
*/
func PrintAlternately_DualChannel(limit int) {
	chanA, chanB := make(chan struct{}), make(chan struct{})

	var wait sync.WaitGroup
	// 两个 goroutine，所以 add(2)
	wait.Add(2)

	go func(limit int) {
		defer wait.Done()

		for i := 1; i <= limit; i++ {
			// 等待 channel A 的信号，有信号时完成一次输出、没有信号时阻塞
			<-chanA

			fmt.Printf("A%d", i)
			chanB <- struct{}{}
		}
	}(limit)

	go func(limit int) {
		defer wait.Done()

		for i := 1; i <= limit; i++ {
			<-chanB

			fmt.Printf("B%d", i)

			// 在最后一轮前，才将控制器交还给 A，最后一轮就直接结束了
			if i < limit {
				chanA <- struct{}{}
			}
		}

	}(limit)

	chanA <- struct{}{}
	wait.Wait()
}

/*
*
单通道版

思考：
- 题目要求：用一个无缓冲通道、两个 goroutine 、严格交替打印
- 两个 goroutine 可以设置为 A goroutine 和 B goroutine
- 一个无缓冲区的通道，记录打印令牌 token，token 中记录打印器、打印进度
- 接收者：
  - goroutine A：接收 channel 事件消息：
  - 若打印器为 A，则进行打印，并将控制权转交给 B
  - 若打印器非 A，则直接将消息填回 channel，让另一个消费者处理
  - goroutine B：接收 channel 事件消息：
  - 若打印器为 A，则进行打印，进度++，到上限时结束，没到上限将控制权转交给 A
  - 若打印器非 A，则直接将消息填回 channel，让另一个消费者处理
*/
func PrintAlternately_SingleChannel(limit int) {
	type printer int
	const (
		aPrinter = +iota
		bPrinter
	)

	// 打印令牌
	type token struct {
		p printer
		i int
	}

	// 无缓冲通道，记录打印令牌
	channel := make(chan token)

	// 等待 A,B 两个 goroutine 完成
	var wait sync.WaitGroup
	wait.Add(2)

	// goroutine A
	go func() {
		defer wait.Done()

		for {
			tk, ok := <-channel
			if !ok {
				return
			}

			// 若不是 A 打印器，则重新塞回去
			if tk.p != aPrinter {
				channel <- tk
				continue
			}

			// 是 A 打印器，则打印
			fmt.Printf("A%d", tk.i)

			// 将打印令牌交给 B
			tk.p = bPrinter
			// 写入通道
			channel <- tk
		}
	}()

	go func() {
		defer wait.Done()

		for {
			tk, ok := <-channel
			if !ok {
				return
			}

			// 若当前不是 B 打印器，则重新塞回去
			if tk.p != bPrinter {
				channel <- tk
				continue
			}

			// 是 B 打印器，则打印
			fmt.Printf("B%d", tk.i)

			// 到达上限，关闭通道，结束 goroutine
			if tk.i >= limit {
				close(channel)
				return
			}

			// 未到达上限，则进度+1，交给 goroutine a
			tk.i++
			tk.p = aPrinter
			channel <- tk
		}
	}()

	channel <- token{p: aPrinter, i: 1}
	wait.Wait()
}
