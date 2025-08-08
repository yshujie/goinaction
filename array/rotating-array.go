package array_test

import "fmt"

/**
✅ 题目：旋转数组

给定一个长度为 n 的数组 arr 和一个整数 k，将数组中的元素向右旋转 k 个位置（原地修改）。

示例
输入：[1, 2, 3, 4, 5, 6, 7], k = 3
输出：[5, 6, 7, 1, 2, 3, 4]

考点
	•	数组索引运算（取模）
	•	原地反转的三步法
	•	时间复杂度 O(n)、空间复杂度 O(1)
*/

/*
*
思考
- 原地修改：说明不要创建新的数组、切片
- 旋转：以 k 为轴，左右旋转，也就是左右替换
- 联想到以数组中心点旋转的题目，那个好处理，从中间点向前遍历，逐个替换
- 再次读题，将数组中的元素向右旋转 k 个位置：
- 等同于从第一个元素开始遍历，将当前元素放置到第 index + k 或 index + k - len(arr) 的位置

k=3
[1, 2, 3, 4, 5, 6, 7]

0+k = 3
1+K = 4
2+K = 5
3+k = 6
4+k = 7-7 = 0
5+k = 8-7 = 1
6+k = 9-7 = 2

这种方法不符合要求：
1. 不是原地反转
2. 空间复杂度是 O(n)
*/
func RotatingArray1(a [7]int, k int) {
	result := a

	for i := 0; i <= len(a)-1; i++ {
		newIndex := i + k
		if i+k >= len(a) {
			newIndex = newIndex - len(a)
		}

		result[newIndex] = a[i]
		fmt.Printf("i:%d; new index: %d; value: %d \n", i, newIndex, result[newIndex])
	}

	fmt.Println(result)
}

/*
*
三步反转法：
- 第一步：反转整个数组，示例：[1, 2, 3, 4, 5, 6, 7] => [7, 6, 5, 4, 3, 2, 1]
- 第二步：反转 [0,k-1]，示例：k = 3 时，[7, 6, 5, 4, 3, 2, 1] => [5, 6, 7, 4, 3, 2, 1]
- 第三步：反转 [k, len(a)-1],示例：k = 3 时，[5, 6, 7, 4, 3, 2, 1] => [5, 6, 7, 1, 2, 3, 4]
*/
func RotatingArray(a []int, k int) {
	// 第一步：反转整个数组
	reverse(a, 0, len(a)-1)
	fmt.Println("a: ", a)

	// 第二步：反转前半段
	reverse(a, 0, k-1)
	fmt.Println("a: ", a)

	// 第三步：反转后半段
	reverse(a, k, len(a)-1)
	fmt.Println("a: ", a)

}

// [1, 2, 3, 4, 5, 6, 7]
func reverse(a []int, start, end int) {
	for start < end {
		a[end], a[start] = a[start], a[end]
		start++
		end--
	}
}
