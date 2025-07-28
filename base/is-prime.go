package base

/**

✅ 题目 3：判断是否为质数
编写函数 IsPrime(n int) bool，使用循环判断一个整数是否为质数。

> 质数的定义：一个大于 1 的正整数，只能被 1 和它本身整除，叫做质数（Prime Number）。

输入示例：
fmt.Println(IsPrime(7))  // true
fmt.Println(IsPrime(10)) // false
*/

func IsPrime(n int) bool {
	// 小于等于 1，都不是质数
	if n <= 1 {
		return false
	}

	for m := n - 1; m > 1; m-- {
		if n%m == 0 {
			return false
		}
	}

	return true
}
