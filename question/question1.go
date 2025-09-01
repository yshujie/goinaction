package question

import (
	"fmt"
	"sync"
)

/**
题目：交替打印数字和字母
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

``
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

/**
思考：
1. goroutine 设置：数字打印 goroutine、字母打印 goroutine
2. channel 设置：数字打印令牌、字母打印令牌
3. 打印权控制：
	a. 起始：主 goroutine 将打印权交给数字打印 goroutine
	b. 交接：
		数字打印 goroutine 打印两次之后，将打印权交接给字母打印 goroutine；
		字母打印 goroutine 打印两次之后，将打印权交接给数字打印 goroutine。
	c. 终止：
		数字打印 goroutine 判断打印值到达 28 后，终止打印，令牌销毁
*/

// 双通道版本的交替打印
func AlternatePrinte() {
	// 创建无缓冲通道
	numberChan := make(chan struct{})
	letterChan := make(chan struct{})

	// 协程计数器
	var wg sync.WaitGroup
	wg.Add(2)

	// 数字打印协程
	go func() {
		defer wg.Done()
		defer close(letterChan)

		i := 1

		for {
			t, ok := <-numberChan
			if !ok { // 数字打印通道已关闭，接收到关闭信号，退出打印
				return
			}

			fmt.Print(i)
			i++
			fmt.Print(i)
			i++

			// 未到达上限，可以交接令牌
			if i < 28 {
				letterChan <- t
				continue
			} else { // 到达上限，结束打印
				return
			}
		}
	}()

	// 字母打印协程
	go func() {
		defer wg.Done()
		defer close(numberChan)

		i := 'A'
		for {
			t, ok := <-letterChan
			if !ok { // 字母打印已通道关闭，接收到关闭信号，退出打印
				return
			}

			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++

			if i < 'Z' { // 未到达上限，交接令牌
				numberChan <- t
				continue
			} else { // 到达上限，结束打印
				return
			}
		}
	}()

	// 给数字打印通道发放令牌，开启打印
	numberChan <- struct{}{}

	// 主协程，等待
	wg.Wait()
}
