package labb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/**
关卡 B：为 sink 的每条写入加局部超时

目标：
- 在 sink 里，每条日志派生 context.WithTimeout(ctx, 300*time.Millisecond)；
- I/O 超过 300ms 视为超时（deadline exceeded），跳过该条但流水线不中断；
- 记得 每次都调用 cancel() 释放定时器资源（不能 defer 在循环里）。

验收：
- 运行时可见：多数条“write ok”，部分条 timeout；
- 全局 5 秒取消后依然优雅退出（看到 [main] graceful done）。并有序退出；
*/

type LogLine struct{ Raw string }

func Hand() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// 使用 context.WithCancel 创建根上下文，用于全局取消；用 defer cancel() 保证释放。
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// 启动一个 5 秒后调用 cancel() 的 goroutine，模拟服务关闭（也可用 time.AfterFunc）。
	// 提示：打印日志 "[main] timed stop -> cancel()"
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
		produce(ctx, lines) // 在 produce 内部：谁生产谁关闭 close(lines)
	}()
	go func() {
		defer wg.Done()
		parse(ctx, lines, parsed) // 处理上游关闭与 ctx 取消，完成后在退出前 close(parsed)
	}()
	go func() {
		defer wg.Done()
		sink(ctx, parsed) // 处理上游关闭与 ctx 取消
	}()

	// 在 main 中“等待退出”：这里可以直接 wg.Wait()。
	wg.Wait()

	log.Printf("[main] graceful done")
}

// produce: 周期性产生日志
func produce(ctx context.Context, out chan<- LogLine) {
	// 创建计时器，每 120 Millisecond 发送一次信号
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()
	// 函数结束前关闭数据通道，遵循“发送端唯一时，谁发送、谁关闭” 的原则，由 producer 关闭 lines 通道
	defer close(out)

	i := 0
	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前协程
			log.Printf("[produce] canceled: %v", ctx.Err())
			return
		case <-ticker.C: // 接收计时器信号
			i++
			out <- LogLine{Raw: fmt.Sprintf("ts=%d level=INFO msg=ok", i)}
		}
	}
}

// parse: 读取 in，做轻量“解析”后发送到 out
func parse(ctx context.Context, in <-chan LogLine, out chan<- LogLine) {
	// 函数结束前关闭通道，遵循“发送者唯一时，谁发送、谁关闭”的原则，由 parser 关闭 parsed 通道
	defer close(out)

	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前进程
			log.Printf("[parse] canceled: %v", ctx.Err())
			return
		case line, ok := <-in:
			if !ok {
				log.Printf("[parse] upstream closed")
				return
			}
			// 这里可以做一些格式校验/过滤，关卡 A 保持直通
			out <- line
		}
	}
}

// sink: 最终消费日志
func sink(ctx context.Context, in <-chan LogLine) {
	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前进程
			log.Printf("[sink] canceled: %v", ctx.Err())
			return
		case line, ok := <-in:
			if !ok {
				log.Printf("[sink] upstream closed")
				return
			}

			// 写入日志，若写入失败，则跳过
			err := writeDownstream(line, ctx)
			if err != nil {
				switch {
				case errors.Is(err, context.DeadlineExceeded):
					log.Printf("[sink] timeout, drop line=%q", line.Raw)
					continue
				case errors.Is(err, context.Canceled):
					log.Printf("[sink] canceled by global ctx")
					return
				default:
					log.Printf("[sink] write error: %v line=%q", err, line.Raw)
					continue
				}
			}

			log.Printf("[sink] write ok line=%q", line.Raw)
		}
	}
}

// writeDownstream 写日志
func writeDownstream(line LogLine, ctx context.Context) error {
	// 设置定时器： 300 毫秒的限时
	itemCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	// 写入后，释放定时器
	defer cancel()

	// 随机一个日志写入时间
	lat := time.Duration(rand.Intn(500)) * time.Millisecond

	select {
	case <-itemCtx.Done(): // 先到达超时限制
		return itemCtx.Err() // ✅ 返回标准错误：DeadlineExceeded/Canceled
	case <-time.After(lat):
		return nil
	}
}
