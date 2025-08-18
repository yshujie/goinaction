package goroutine_test

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/**
✅ 题目：用 select 实现“火箭发射：倒计时与中止”
实现一个简化版的火箭发射控制台程序。程序会进行 逐秒倒计时，在倒计时结束时打印“Liftoff!” 并发射；
但在倒计时期间，用户可以按下 Enter 来中止发射，程序需立刻停止倒计时并打印“Launch aborted!”。

要求：
- 每秒打印一次倒计时剩余秒数（例如 T-10, T-9, …, T-1）。
- 倒计时尚未结束时，用户在终端按下 Enter（回车）：立即中止并打印 “Launch aborted!”
- 如果倒计时完成且未被中止：打印 “Liftoff!”
- 必须使用 select 来同时等待“用户中止事件”和“计时事件”。
- 需要 优雅退出：不得有 goroutine 泄露；用到的 Ticker 需 Stop()。
- 不得用 time.Sleep 循环代替 Ticker/Timer；不得通过 busy loop 轮询输入。

边界与鲁棒性：
- n <= 0：直接打印 Liftoff!（或不倒计时直接发射），程序正常退出。
- 输入读取应非阻塞地集成到 select（例如通过单独 goroutine 发信号或使用带 context 的读）。
- 在极端情况下（用户在最后 1 秒内按下 Enter），以 “谁先到达 select 就处理谁” 为准：
- 如果先收到中止信号 ⇒ 中止；
- 如果先到达倒计时结束 ⇒ 发射。
*/

func RocketLaunch(n int) {
	if n <= 0 {
		launch()
		return
	}

	// 1. 回车终止：一次性信号，缓存为 1，避免发送端阻塞
	abort := make(chan struct{}, 1)
	go func() {
		fmt.Println("Press Enter to abort launch.")

		// goroutine 阻塞，等待输入回车符
		_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		select {
		case abort <- struct{}{}: // 如果放得进去，就放
		default: // 放不进去（说明已经有一个信号了），什么也不做，直接返回
		}
	}()

	// 2. 逐秒心跳 + n 秒“闸门”
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// 延迟触发器，单独掌控发射时刻
	deadline := time.After(time.Duration(n) * time.Second)

	// 3. 单循环，三路竞争：中止 / 到点发射 / 每秒打印
	remeining := n
	for {
		select {
		case <-abort: // 收到终止信号，终止发射
			abortLaunch()
			return
		case <-deadline: // 到达发射时间，进行发送
			launch()
			return
		case <-ticker.C: // 倒计时
			fmt.Println("倒计时：", n)
			remeining--
		}
	}
}

func launch() {
	fmt.Println("Liftoff!")
}

func abortLaunch() {
	fmt.Println("Launch aborted!")
}
