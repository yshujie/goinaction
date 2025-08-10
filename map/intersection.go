package map_test

/**
✅ 题目：两个数组的交集
给定两个整数数组 nums1 和 nums2，返回它们的交集（去重）。

示例：
输入：nums1 = [1, 2, 2, 1], nums2 = [2, 2]
输出：[2]

考点：
	•	map[int]bool 存储
	•	去重处理
	•	时间复杂度 O(n)
*/

/**
思路：
- 求两个切片的交集（去重），就是找某些元素即存在于 nums1，也存在于 nums2
- 可创建一个 map[item]int, 记录每个 item 出现次数（在一个切片中出现多次，算一次）
- 最后遍历 map[item]int，获取 int = 2 的 items 集合
*/

func Intersection(numss ...[]int) (result []int) {
	// 创建统计表
	cntTable := make(map[int]int)
	for _, nums := range numss {
		tmpTable := make(map[int]struct{})
		for _, num := range nums {
			// 第一次出现某个元素，则记录该元素出现
			if _, ok := tmpTable[num]; !ok {
				tmpTable[num] = struct{}{}
				cntTable[num]++
			}
		}
	}

	// 统计交集元素
	for num, cnt := range cntTable {
		if cnt == len(numss) {
			result = append(result, num)
		}
	}

	return result
}
