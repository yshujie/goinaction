package base

/**
✅ 题目四：翻转一个切片（原地修改）
写一个函数，反转一个 int 类型的切片
输入: []int{1, 2, 3, 4}
输出: []int{4, 3, 2, 1}


✍ 要求
- 原地反转，不能新建切片
- 时间复杂度 O(n)
*/

func ReverseSlice(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return s
}
