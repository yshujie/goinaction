package map_test

/*
*
✅ 找出出现次数最多的元素
给定一个整数切片，找出出现次数最多的元素和它的次数（支持多个并列最多）。

示例：
输入：[1, 3, 2, 3, 4, 3, 2, 2]
输出：元素 2 出现 3 次；元素 3 出现 3 次

考点：
  - map[int]int 统计频率
  - 遍历找最大值
  - 多结果处理
*/
func FindMostFrequentItem(i []int) (maxCount int, items []int) {
	if len(i) == 0 {
		return 0, items
	}

	// 使用 map[int]int 做统计表
	countTable := make(map[int]int)
	for _, item := range i {
		countTable[item]++
	}

	// 查找出现的最大次数
	for _, cnt := range countTable {
		if cnt > maxCount {
			maxCount = cnt
		}
	}

	// 根据最大次数查找元素
	items = items[:0]
	for item, cnt := range countTable {
		if cnt == maxCount {
			items = append(items, item)
		}
	}

	return maxCount, items
}
