package slice_test

/**
✅ 题目：去除切片中的重复元素（保持原顺序）
输入一个 []int，去重并保持第一次出现的顺序。

示例：
输入：[1, 2, 3, 2, 1, 4]
输出：[1, 2, 3, 4]

考点：
	•	map[int]bool 去重
	•	切片 append 顺序控制
	•	切片是引用语义（输出与原切片的关系）
*/

/**
思路：
- 读写双指针 + 就地去重法（不创建新的切片，而是通过更新原切片的底层数组来实习）
- 读指针，从左向右遍历切片
- 写指针，遇到第一次出现的元素时进行更新切片数据
- map 记录已经写入过的元素
*/

func RemoveDuplicates[T comparable](i []T) []T {
	if len(i) <= 1 {
		return i
	}

	// 元素标记（存储元素是否被记录）
	seen := map[T]struct{}{}
	// 写指针
	writerIndex := 0

	// 读切片，更新底层数组
	for _, v := range i {
		// v 已经出现过，跳过
		if _, ok := seen[v]; ok {
			continue
		}

		// v 第一次出现，则标记元素、更新底层数组、更新写指针
		seen[v] = struct{}{}
		i[writerIndex] = v
		writerIndex++
	}

	// 清理切片尾巴
	clear(i[writerIndex:])

	return i[:writerIndex]
}
