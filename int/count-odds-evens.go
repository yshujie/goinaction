package base

import "fmt"

/*
*
✅ 题目 1：统计奇数和偶数的个数
编写函数 CountOddsEvens(nums []int)，接收一个整数切片，打印奇数与偶数的数量。

输入示例：
nums := []int{1, 2, 3, 4, 5, 6}
CountOddsEvens(nums)

输出示例：
奇数个数: 3
偶数个数: 3
*/
func CountOddsEvens(nums []int) {
	oddCnt, evenCnt := 0, 0

	for _, x := range nums {
		if isOdd(x) {
			oddCnt++
		} else {
			evenCnt++
		}
	}

	fmt.Printf("奇数个数：%d \n", oddCnt)
	fmt.Printf("偶数个数：%d \n", evenCnt)
}

// isOdd 是否为奇数
func isOdd(x int) bool {
	return x%2 != 0
}
