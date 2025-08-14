package goroutine_test

import (
	"fmt"
	"time"
)

func CalcFibNum(x int) {
	// 启动协程，显示跳动的光标，表示程序执行中
	go spinner(10 * time.Millisecond)

	// 计算斐波那契数
	fibNum := fib(x)
	fmt.Printf("\r %d 's fib number is %d \n", x, fibNum)
}

// Spinner 显示字符串，表示程序存活中
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|\` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// Fib 计算斐波那契数
func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
