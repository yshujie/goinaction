package base

import "fmt"

// 请使用 Go 语言编写函数 PrintMultiplicationTable(n int)，打印 n × n 的乘法表。
// 示例输入：n = 3
// 1*1=1
// 1*2=2 2*2=4
// 1*3=3 2*3=6 3*3=9
// 👉 要求使用嵌套循环实现，格式控制对齐不是必须，但输出内容要正确。

// PrintMultiplicationTable 打印 n*n 乘法表
// row 控制行， x，y 控制乘法运算
func PrintMultiplicationTable(n int) {
	if n <= 0 {
		fmt.Printf("error params n, %d is less than 0 \n", n)
		return
	}

	for row := 1; row <= n; row++ {
		for x, y := 1, row; x <= y; x++ {
			if x < y {
				fmt.Printf("%d*%d=%d  ", x, y, x*y)
			} else if x == y {
				fmt.Printf("%d*%d=%d\n", x, y, x*y)
			}
		}
	}
}

// PrintMultiplicationTable 打印 n*n 乘法表
// 使用 row + col 控制，乘法表的参数只是 列 * 行
func PrintMultiplicationTable2(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			if row == col {
				fmt.Printf("%d*%d=%d \n", col, row, col*row)
			} else {
				fmt.Printf("%d*%d=%d ", col, row, col*row)
			}
		}
	}
}

// PrintMultiplicationTable3 打印 n*n 乘法表
// 使用 col + row 控制，并优化输出
func PrintMultiplicationTable3(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Printf("%d*%d=%-2d ", col, row, col*row)
		}
		fmt.Println()
	}
}
