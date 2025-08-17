package array_test

/**
✅ 题目：判断数组是否单调递增
实现一个函数，判断一个数组是否是单调递增或不减的。

示例
输入：[1, 2, 2, 3] → 输出：true
输入：[1, 3, 2] → 输出：false

考点
	•	顺序遍历 + 相邻元素比较
	•	边界情况（空数组、单元素）
*/

/*
*
思考：
- 判断数组是否为单调递增，则可以从第二个元素开始遍历数组
- 若当前元素 < 上一个元素，则返回 false
- 若当前元素 >= 上一个元素，则 continue

- 对于临界值的思考：
- 若数组元素个数 <= 1， 则直接返回 true
- 遍历时，要遍历到数组最后一个元素
*/
func IsIncreasing(a []int) bool {
	if len(a) <= 1 {
		return true
	}

	for i := 1; i <= len(a)-1; i++ {
		if a[i] < a[i-1] {
			return false
		}
	}

	return true
}
