package array_test

/**
✅ 题目：反转数组
使用数组指针作为参数，反转数组元素

要求：
1. reverse 函数的参数类型为 *[N]T，其中 N 是数组长度，T 是数组元素类型。
2. 函数需要原地反转数组中的元素（不分配新内存）。

函数签名示例：
func reverse(arr *[N]int)

示例运行：
输入: [0 1 2 3 4 5]
输出: [5 4 3 2 1 0]

注意：
- 不能直接用 append 或 slice 操作，要通过数组下标交换实现。
- 需要用指针解引用访问数组元素，例如：
  (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
*/

func Reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[j], a[i] = a[i], a[j]
	}
}
