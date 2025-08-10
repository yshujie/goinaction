package slice_test

/**
✅ 题目：删除切片中指定值的所有元素
输入一个切片 s 和一个值 val，删除切片中所有等于 val 的元素，返回新的切片。

示例：
输入：s = [1, 2, 3, 2, 4], val = 2
输出：[1, 3, 4]

考点：
	•	切片“原地”过滤技巧：s = s[:0] 后 append
	•	遍历时的写入覆盖技巧
	•	不额外开辟新切片
*/

/**
思路：
- 使用 读写双指针 + 原地更新法
- 从左向右遍历切片，判断每一个元素是否为值得值 val
	若是，则跳过当前值
	若不是，则写入读取的值，并更新写入指针
*/

func RemoveSilceValue[T comparable](s []T, val T) []T {
	if len(s) == 0 {
		return s
	}

	// 写指针
	w := 0

	// 读取切片元素
	for _, currVal := range s {
		// 遇到需要被过滤的值，则遗弃
		if currVal == val {
			continue
		}

		// 非过滤值，则写入切片，更新写指针进度
		s[w] = currVal
		w++
	}

	// 切片尾部更新，优化 GC
	clear(s[w:])

	return s[:w:w]
}
