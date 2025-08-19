package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：给 10 个协程，如何打印出升序的数字
*/
/**
思考：
- 这是一个多协程访问公共变量的问题：
- 10 个协程相当于 10 个 worker
- 打印升序数字，相当于 10 个 worker 访问公共的内存空间，获取打印权限，进行打印
- 进一步说明，协程间是竞争关系

解法：
- 创建无缓冲 channel ，存储打印令牌，哪个协程获取到令牌，则该协程获取数字，进行打印
*/

// 随机 worker 模式
func PrintAscendingWithRundomWorker(max, workerCnt int) {
	if max <= 0 || workerCnt <= 0 {
		return
	}

	// 单槽令牌，避免启动竞态
	tokenChan := make(chan int, 1)

	// 设置等待组
	var wg sync.WaitGroup
	wg.Add(workerCnt)

	// 创建 worker 进行打印
	for i := 0; i < workerCnt; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				num, ok := <-tokenChan
				if !ok {
					return
				}

				fmt.Printf("[work: %d] print: %d \n", id, num)
				num++

				if num <= max {
					tokenChan <- num
				} else {
					close(tokenChan)
					return
				}
			}
		}(i)
	}

	// 开启打印
	tokenChan <- 0

	wg.Wait()
}

// 顺序 worker 模式，严格轮转“接力棒”模型（固定 worker 顺序）
// 每个 worker 分配自己的 channel，轮到 worker 执行时，向对应的 channel 中下发通知
func PrintAscendingWithOrderWorker(max, workerCnt int) {
	if max < 0 || workerCnt <= 0 {
		return
	}

	// 接力棒
	type Baton struct{ num int }
	// 创建 batons 切片，存储 baton 通道
	batonChans := make([]chan Baton, workerCnt)
	for i := 0; i < workerCnt; i++ {
		batonChans[i] = make(chan Baton)
	}

	// 关闭广播通道
	done := make(chan struct{})

	var once sync.Once

	// 设置等待组
	var wg sync.WaitGroup
	wg.Add(workerCnt)

	// 启动多个 worker 进行打印
	for i := 0; i < workerCnt; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				select {
				case <-done: // 收到关闭广播，停止程序
					return
				case baton, ok := <-batonChans[id]:
					if !ok {
						return
					}

					// 打印
					fmt.Printf("[worker: %d] %d \n", id, baton.num)
					// 计数
					baton.num++

					// 已到超过大值则停止，并发布关闭广播、关闭所有 worker 通道
					if baton.num > max {
						once.Do(func() { close(done) })
						return
					}

					// 未到达最大值，则接力给下一个 worker
					// 发送前复查 done，防止外部取消竞态
					select {
					case <-done:
						return
					case batonChans[nextId(id, workerCnt)] <- baton:
					}
				}
			}
		}(i)
	}

	// 开始信号
	batonChans[0] <- Baton{num: 0}

	wg.Wait()
}

func nextId(id, workerCnt int) int {
	return (id + 1) % workerCnt
}
