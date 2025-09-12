package array_test

import "sync"

/**
题目：golang 编写代码实现一个 channel 数组求和，输入为 chan int 数组，输出为int。

思考：
- 参数是 chan int 数组，也就是会有多个 channel
- 将所有 channel 的 int 元素汇集到一个 channel 中，最后再通过这个汇集后的 channel 求和
*/

func SumAll(chs []<-chan int) (result int) {
	// 创建有缓冲的 sum channel 通道
	sumCh := make(chan int, 100)

	// 设置协程等待器
	var wg sync.WaitGroup

	// 读取参数中的每一个通道中的元素
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()

			// 将通道内的所有元素发送到 sum channel
			for v := range ch {
				sumCh <- v
			}
		}(ch)
	}

	// 监控协程情况，做完元素搬迁后关闭 sum channel 通道
	go func() {
		wg.Wait()
		close(sumCh)
	}()

	// 将所有的元素累加
	for v := range sumCh {
		result += v
	}

	return result
}
