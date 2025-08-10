package slice_test

/**
✅ 题目：合并两个有序切片并保持有序
输入两个升序切片 a 和 b，合并为一个升序切片。

示例：
输入：a = [1, 3, 5], b = [2, 4, 6]
输出：[1, 2, 3, 4, 5, 6]

考点：
	•	双指针遍历
	•	切片 append
	•	时间复杂度 O(n)
*/

/**
思路：
- 将两个切片按照升序进行合并，这是“线性归并”问题
- 另外，由于两个原切片本身就是升序切片，所以可以利用这个特性，减少一部分运算
- 创建 cap = len(a) + len(b) 的新切片
- 遍历 a 和 b 切片，每次获取 a 或 b 切片的最小值进行比较，谁小用谁
- 当 a / b 有其一用尽时，直接将有剩余的切片追加进新切片就好了
*/

func MergeSortedInts(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	// a / b 仍有剩余，则在 a / b 切片中寻找最小的元素，写入 result 中
	j, k := 0, 0
	for j < len(a) && k < len(b) {
		// 获取 a / b 中较小的一个
		if a[j] < b[k] {
			result = append(result, a[j])
			j++
		} else {
			result = append(result, b[k])
			k++
		}
	}

	// a / b 中有一个被用完了，则将另一个全部追加入 result 中
	if j < len(a) {
		result = append(result, a[j:]...)
	} else if k < len(b) {
		result = append(result, b[k:]...)
	}

	return result
}
