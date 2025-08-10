package array_test

/**
✅ 题目：查找数组中第二大的元素
实现一个函数，返回数组中第二大的数字（要求一次遍历完成）。

示例
输入：[3, 1, 4, 5, 2]
输出：4

考点
	•	一次遍历维护两个变量 max1、max2
	•	边界值判断（数组长度 < 2）
*/

func FindSecondItem(a []int) int {
	if len(a) <= 1 {
		return 0
	}

	// 将初始值设为第一个元素
	max1, max2 := a[0], a[0]

	// 从第二个元素开始遍历
	for _, v := range a[1:] {
		if v > max1 { // 发现比 max1 更大的元素，则更新 max1，并将原 max1 设置为 max2
			max2 = max1
			max1 = v
		} else if max1 > v && max2 == a[0] { // 发现元素比 max1 小，且 max2 为初始值时，则设置为 max2
			max2 = v
		} else if max1 > v && v > max2 { // 发现元素比 max1 小，且比 max2 大，则设置为第二大值
			max2 = v
		}
	}

	return max2
}
