package goroutine_test

import (
	"fmt"
	"sync"
)

/**
✅ 题目：A5 单向通道
实现一个使用单向 channel 限制读写权限的程序，生产者函数只负责写数据，消费者函数只负责读数据，保证函数签名限制了其对 channel 的访问权限。

函数格式：
func producer(ch chan<- int)
func consumer(ch <-chan int)

输入：
无

输出：
生产者发送的数据；消费者接收并打印。

考点：
- 单向 channel 的声明与使用
- API 设计中的访问权限限制
- 多函数协作传递 channel

*/

func ProducerConsumerCtr() {
	// 创建有缓存的通道
	ch := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)

	go consumer(ch, &wg)

	// 主进程等待
	wg.Wait()
}

// 生产者，向 channel 中写数据
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // 由生产者关闭通道

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range s {
		fmt.Println("生成数据：", v)
		ch <- v
	}
}

// 消费者，从 channel 中读数据，并输出
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range ch {
		fmt.Println("读取数据：", item)
	}
}
