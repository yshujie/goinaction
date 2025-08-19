package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：交替打印数字和字母
使用两个 goroutine 交替打印数字和字母，如：1A2B3C...

注意：
- A 打印在先，B 打印在后
- 只能一个通道传令牌；严禁向“自己负责接收”的通道再发送。
- 用 WaitGroup 收尾；退出无悬挂、无死锁；go test -race 通过。
*/

/**
思考：
- 两个 goroutine ，一个打印数字、一个打印字母
- 顺序打印，即严格控制 goroutine 的执行顺序，使用无缓冲的 channel 传递 token
- 结束标志：到达字母 Z 结束
*/

type mode int

const (
	numberMode = +iota
	letterMode
)

type token struct {
	mode mode
}

// 单通道版本
// 使用一个 channel 控制令牌
func AlternatePrintingSingleChan() {
	// 创建无缓冲的 channel
	tokenChan := make(chan token)

	// 创建等待
	var wait sync.WaitGroup
	wait.Add(2)

	// 创建数字打印 goroutine
	go func() {
		defer wait.Done()

		i := 1
		for {
			t, ok := <-tokenChan
			if !ok {
				return
			}
			if t.mode != numberMode {
				tokenChan <- t
				continue
			}

			fmt.Print(i)
			i++

			tokenChan <- token{mode: letterMode}
		}
	}()

	// 创建字母打印 goroutine
	go func() {
		defer wait.Done()

		var l rune = 'A' // 'A' 是 rune 类型字面量
		var last rune = 'Z'
		for {
			t, ok := <-tokenChan
			if !ok {
				return
			}
			if t.mode != letterMode {
				tokenChan <- t
				continue
			}

			// 打印字母
			fmt.Print(string(l))

			// 控制 token
			if l < last {
				l++
				tokenChan <- token{mode: numberMode}
			} else {
				close(tokenChan)
				return
			}
		}
	}()

	// 向 token 中
	tokenChan <- token{mode: numberMode}

	// 主 goroutine 等待
	wait.Wait()
}

// 双通道版本
// 使用两个 channel ，分别控制数字打印令牌、字母打印令牌
func AlternatePrintingDualChan() {
	// 创建两个无缓冲通道，分别控制字母打印、数字打印
	numberChan := make(chan struct{})
	letterChan := make(chan struct{})

	// 创建等待
	var wait sync.WaitGroup
	wait.Add(2)

	// 数字打印 goroutine
	go func() {
		defer wait.Done()

		i := 1
		for {
			_, ok := <-numberChan
			// number channner 被关闭后触发
			if !ok {
				return
			}

			// 打印数字、并计数
			fmt.Print(i)
			i++

			// 交接令牌，触发字母打印
			letterChan <- struct{}{}
		}
	}()

	// 字母打印 goroutine
	go func() {
		defer wait.Done()

		l := 'A'
		for {
			_, ok := <-letterChan
			// letter channel 被关闭后触发
			if !ok {
				return
			}

			// 打印字母
			fmt.Print(string(l))
			l++

			if l <= 'Z' {
				// 交接令牌，触发数字打印
				numberChan <- struct{}{}
			} else {
				// 结束打印，关闭通道
				close(letterChan)
				close(numberChan)
				return
			}
		}

	}()

	// 向 number channel 中传递打印令牌
	numberChan <- struct{}{}

	wait.Wait()
}
