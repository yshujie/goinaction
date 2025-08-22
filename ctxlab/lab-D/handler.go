package labd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

/**
关卡 C：OS 信号优雅关停 + traceID 传递

目标：
- 捕获 SIGINT/SIGTERM → cancel() → 全线优雅退出；
- 为每条日志生成并传递 traceID（只作为请求域只读元数据展示链路），不要用 WithValue 传业务大对象；
- 保留 C 关的 errgroup 编排与单条超时。

任务：
- 使用 signal.Notify 监听 SIGINT/SIGTERM；收到即 cancel()。
-在 producer 为每条日志生成 traceID（简易随机字符串），用 context.WithValue 传给下游，仅用于日志打印链路追踪。
- 限制：不要用 WithValue 传业务大对象；只传小型只读元数据（如 traceID）。

验收：
Ctrl+C 时看到协程逐步退出；日志打印含 traceID（如 [trace=abc123] write ok）。

*/

type LogLine struct {
	Raw     string
	TraceID string // ✅ 路线A：把 traceID 放进业务结构体

}

func Hand() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	rand.Seed(time.Now().UnixNano())

	// 使用 context.WithCancel 创建根上下文，用于全局取消；用 defer cancel() 保证释放。
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// 监听 OS 信号（SIGINT/SIGTERM），收到后 log 一行并调用 cancel()
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigch
		log.Printf("[main] caught signal: %s -> cancel()", s.String())
		cancelFunc()
	}()

	// 启动一个 5 秒后调用 cancel() 的 goroutine，模拟服务关闭（也可用 time.AfterFunc）。
	// 提示：打印日志 "[main] timed stop -> cancel()"
	time.AfterFunc(10*time.Second, func() {
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
			raw := randomLine(i)
			// 生成一个 8 位 traceID（建议调用 randTrace()）
			tid := randTrace()

			// 把 traceID 写进 LogLine；
			out <- LogLine{Raw: raw, TraceID: tid}
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

			// 首次命中 "FATAL" 则返回错误（例如 errors.New("fatal log encountered")）
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

			// 打印成功日志时带上 traceID
			log.Printf("[sink] write ok [trace=%s] line=%q", line.TraceID, line.Raw)
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

// 生成 8 位 traceID
func randTrace() string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
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
