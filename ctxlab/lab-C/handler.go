package labc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

/**
关卡 C：errgroup.WithContext 错误收敛

目标：
- 用 errgroup.WithContext 替换 WaitGroup 进行统一编排。
- parse 一旦遇到包含 "FATAL" 的日志行，返回 error 触发全线取消。
- 其他协程不再返回二次错误（例如 context.Canceled），而是识别为级联取消后“正常收尾返回 nil”，避免覆盖首个真实错误。
- 继续保留关卡 B 的“单条 I/O 超时”行为（300ms 丢弃该条，流水线不中断）。

验收：
- 运行后，你会看到持续处理；当产生 FATAL 行（5% 概率），立刻全线停机：parse 报错；produce/sink 迅速打印“canceled/closed”并退出；main 输出 exit with error: fatal log encountered。
- 若 5 秒内没命中 FATAL，也会因 cancel() 优雅退出（这是正常分支）。

*/

type LogLine struct{ Raw string }

func Hand() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	rand.Seed(time.Now().UnixNano())

	// 使用 context.WithCancel 创建根上下文，用于全局取消；用 defer cancel() 保证释放。
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// 启动一个 5 秒后调用 cancel() 的 goroutine，模拟服务关闭（也可用 time.AfterFunc）。
	// 提示：打印日志 "[main] timed stop -> cancel()"
	time.AfterFunc(5*time.Second, func() {
		log.Printf("[main] timed stop -> cancel()")
		cancelFunc()
	})

	// 用 run(root) 返回的 error 决定主流程日志：
	// 有错打印 "[main] exit with error: %v"；否则 "[main] graceful done"
	if err := run(ctx); err != nil {
		log.Printf("[main] exit with error: %v", err)
	} else {
		log.Printf("[main] graceful done")
	}
}

func run(rootCtx context.Context) error {
	rand.Seed(time.Now().UnixNano())

	g, ctx := errgroup.WithContext(rootCtx)

	// 管道：生产者 -> 解析 -> 下游
	lines := make(chan LogLine, 32)
	parsed := make(chan LogLine, 32)

	// producer：唯一写 lines，负责关闭
	g.Go(func() error {
		defer close(lines)         // 唯一生产者：谁生产，谁关闭
		return produce(ctx, lines) // 确保 produce 返回 error（而不是直接退出）
	})

	// parser：唯一写 parsed，负责关闭；命中 FATAL 要返回 error 触发取消
	g.Go(func() error {
		defer close(parsed)              // 唯一生产者：谁生产，谁关闭
		return parse(ctx, lines, parsed) // 命中 FATAL 要返回 error 触发取消
	})

	// sink：消费 parsed；尊重 ctx.Done()；注意“被级联取消时返回 nil”
	g.Go(func() error {
		return sink(ctx, parsed) // 处理 item 超时；ctx 级联取消时返回 nil
	})

	return g.Wait()
}

// produce: 周期性产生日志
func produce(ctx context.Context, out chan<- LogLine) error {
	// 创建计时器，每 120 Millisecond 发送一次信号
	ticker := time.NewTicker(120 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前协程
			log.Printf("[produce] canceled: %v", ctx.Err())
			return nil
		case <-ticker.C: // 接收计时器信号
			i++
			out <- LogLine{Raw: randomLine(i)}
		}
	}
}

// parse: 读取 in，做轻量“解析”后发送到 out
func parse(ctx context.Context, in <-chan LogLine, out chan<- LogLine) error {
	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前进程
			log.Printf("[parse] canceled: %v", ctx.Err())
			return nil
		case line, ok := <-in:
			if !ok {
				log.Printf("[parse] upstream closed")
				return nil
			}

			// 命中 "FATAL" 则返回错误（例如 errors.New("fatal log encountered")）
			if containsFatal(line.Raw) {
				log.Printf("[parse] hit FATAL: %q", line.Raw)
				return errors.New("fatal log encountered")
			}

			// 将日志写入输出通道
			out <- line
		}
	}
}

// sink: 最终消费日志
func sink(ctx context.Context, in <-chan LogLine) error {
	for {
		select {
		case <-ctx.Done(): // 接收关停广播，结束当前进程
			log.Printf("[sink] canceled: %v", ctx.Err())
			return nil
		case line, ok := <-in:
			if !ok {
				log.Printf("[sink] upstream closed")
				return nil
			}

			// 写入日志，若写入失败，则跳过
			err := writeDownstream(ctx, line)
			if err != nil {
				switch {
				case errors.Is(err, context.DeadlineExceeded):
					log.Printf("[sink] timeout, drop line=%q", line.Raw)
					continue
				case errors.Is(err, context.Canceled):
					log.Printf("[sink] canceled by global ctx")
					return nil
				default:
					log.Printf("[sink] write error: %v line=%q", err, line.Raw)
					continue
				}
			}

			log.Printf("[sink] write ok line=%q", line.Raw)
		}
	}
}

// randomLine 随机产生日志信息
func randomLine(i int) string {
	// 5% 概率注入 FATAL
	if rand.Intn(20) == 0 {
		return fmt.Sprintf("ts=%d level=FATAL msg=boom", i)
	}
	return fmt.Sprintf("ts=%d level=INFO msg=ok", i)
}

func containsFatal(s string) bool {
	return len(s) >= 5 && (indexOf(s, "FATAL") >= 0)
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

// writeDownstream 写日志
func writeDownstream(ctx context.Context, line LogLine) error {
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
