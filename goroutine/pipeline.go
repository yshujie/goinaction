package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：“计数 → 求平方 → 打印”的流水线
用三个 goroutine 串联完成“计数 → 求平方 → 打印”的流水线

你需要实现一个最小可用的 goroutine + channel 流水线：
- goroutine A（counter）按序产生整数 0,1,2,...,N-1；
- goroutine B（square）接收整数，计算平方；
- goroutine C（printer）接收平方结果并输出。

要求：用 channel 串联 三个阶段，并正确处理 关闭通道、退出时序，避免死锁。
- 创建三个 goroutine：counter、square、printer，用两个通道连接：
  - ch1: counter → square
  - ch2: square → printer
- counter 产生 0..N-1 的整数并发送到 ch1，然后 关闭 ch1。
- square 从 ch1 读取，计算平方后发送到 ch2；当 ch1 读尽后，关闭 ch2。
- printer 持续从 ch2 读取并打印到标准输出（逐行打印）。
- main 必须 等待所有工作完成后退出（不能靠 time.Sleep 盲等）。
- 程序接收一个命令行参数 -n（默认 10），表示要生成的整数个数。
- 不允许使用全局变量作为数据通道或同步手段。
*/

func Pipeline(n int) {
	// 判断 n 的合法性
	if n <= 0 {
		return
	}

	// counter goroutine，遍历整数，发送带 ch1 通道
	ch1 := conter(n)

	// square goroutine，计算平方，发送到 ch2 通道
	ch2 := square(ch1)

	// printer goroutine，读取 ch2 通道内容，进行打印
	var w sync.WaitGroup
	printer(ch2, &w)

	w.Wait()
}

// conter 函数，将 0～n-1 持续输入到 channel 中
func conter(n int) <-chan int {
	// 创建无缓冲通道
	out := make(chan int)

	// 创建 sub goroutine，持续向通道内注入值
	go func() {
		defer close(out)

		for i := 0; i <= n-1; i++ {
			out <- i
		}
	}()

	return out
}

// square 函数，对输入通道元素做平方计算
func square(in <-chan int) <-chan int {
	out := make(chan int)

	// 创建 sub goroutine，读取 in 通道内数据，做平方计算后写入 out 通道
	go func() {
		defer close(out)

		for x := range in {
			out <- x * x
		}
	}()

	return out
}

// printer 函数，打印结果
func printer(ch <-chan int, wait *sync.WaitGroup) {
	wait.Add(1)
	// 启动 sub goroutine，持续读取数据，并输出
	go func() {
		defer wait.Done()

		for x := range ch {
			fmt.Println(x)
		}
	}()
}
