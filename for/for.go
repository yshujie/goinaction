package for_test

import (
	"fmt"
)

// ContinuousPrint 连续打印
// 根据参数，从 start 打印到 end
func ContinuousPrint(start, end int) {
	for index := start; index <= end; index++ {
		fmt.Printf("index: %d \n", index)
	}
}

// ContinuousPrint2 连续打印
// 直接使用 start， end 参数进行循环打印
func ContinuousPrint2(start, end int) {
	for ; start <= end; start++ {
		fmt.Printf("start: %d \n", start)
	}
}

// ContinuousPrint3 连续打印
// 声明多变量的形式，进行循环打印
func ContinuousPrint3(start, end int) {
	for x, y := start, end; x <= y; x++ {
		fmt.Printf("x: %d, y: %d \n", x, y)
	}
}

// ContinousPrint4 连续打印
// 直接判断
func ContinuousPrint4(start, end int) {
	for start <= end {
		fmt.Printf("index: %d; left lenght: %d \n", start, end-start)
		start++
	}
}

// ContinuousPrint5 连续打印
// 死循环形式
func ContinuousPrint5(start, end int) {
	for {
		index := start
		fmt.Printf("index: %d \n", index)

		if start < end {
			start++
		} else {
			break
		}
	}
}

// ContinuousPrint6 连续打印
// for range 形式
func ContinuousPrint6(start, end int) {
	sl := []int{}

	for item := start; item <= end; item++ {
		sl = append(sl, item)
	}

	for index, value := range sl {
		fmt.Printf("index: %d ; value: %d \n", index, value)
	}
}
