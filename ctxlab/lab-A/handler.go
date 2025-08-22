package laba

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

/**
关卡 A：全局取消 + 基础流水线

目标：
- 用 context.WithCancel 产生全局取消信号；
- 搭建三段流水线：producer -> parse -> sink；
- 谁生产谁关闭通道（一次性关闭广播），下游收到关闭后自然退出；
- 主 goroutine Wait()，最终优雅退出（打印 [main] graceful done）。

验收标准：
- 程序启动后持续输出 sink write ok；
- 5 秒后自动 cancel()，各阶段打印取消或上游关闭日志并有序退出；
- 程序结尾打印 [main] graceful done。
*/

type LogLine struct{ Raw string }

func Hand() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// 5 秒后模拟关停：先打日志，再取消；不要用 Fatal
	time.AfterFunc(5*time.Second, func() {
		log.Printf("[main] timed stop -> cancel()")
		cancelFunc()
	})

	lines := make(chan LogLine, 32)
	parsed := make(chan LogLine, 32)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		produce(ctx, lines)
	}()
	go func() {
		defer wg.Done()
		parse(ctx, lines, parsed)
	}()
	go func() {
		defer wg.Done()
		sink(ctx, parsed)
	}()

	wg.Wait()
	log.Printf("[main] graceful done")
}

// 周期性产生日志
func produce(ctx context.Context, out chan<- LogLine) {
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()
	defer close(out) // ✅ 谁生产谁关闭；无论何种 return 都能保证关闭

	i := 0
	for {
		select {
		case <-ctx.Done():
			log.Printf("[produce] canceled: %v", ctx.Err())
			return
		case <-ticker.C:
			i++
			out <- LogLine{Raw: fmt.Sprintf("ts=%d level=INFO msg=ok", i)}
		}
	}
}

// 读取 in，做轻量“解析”后发送到 out
func parse(ctx context.Context, in <-chan LogLine, out chan<- LogLine) {
	defer close(out) // ✅ 这里是唯一写 out 的地方，安全关闭

	for {
		select {
		case <-ctx.Done():
			log.Printf("[parse] canceled: %v", ctx.Err())
			return
		case line, ok := <-in:
			if !ok {
				log.Printf("[parse] upstream closed")
				return
			}
			out <- line
		}
	}
}

// 最终消费日志
func sink(ctx context.Context, in <-chan LogLine) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("[sink] canceled: %v", ctx.Err())
			return
		case line, ok := <-in:
			if !ok {
				log.Printf("[sink] upstream closed")
				return
			}
			time.Sleep(50 * time.Millisecond) // 模拟处理
			log.Printf("[sink] write ok line=%q", line.Raw)
		}
	}
}
