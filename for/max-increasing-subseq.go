package for_test

/*
✅ 题目 5：找出切片中最大连续递增子序列的长度
编写函数 MaxIncreasingSubseq(nums []int) int，找出切片中最长的连续递增子序列长度。

输入示例：
nums := []int{1, 2, 2, 3, 4, 1, 2, 3, 4, 5}
fmt.Println(MaxIncreasingSubseq(nums))

输出：
5
*/

// 思路：最大联系递增子序列：
// 切片长度 <=1 的，直接是切片长度；
// 切片长度 > 1 的，遍历切片，窗口为 2 ，每次递增 1 位，窗口内 后-前 = 1，则当前递增序列长度 + 1，否则，当前的证需要长度置为 0
func MaxIncreasingSubseq(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	maxLen := 1
	currentLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentLen++
		} else {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 1
		}
	}

	// 注意结尾可能是最大递增序列
	if currentLen > maxLen {
		maxLen = currentLen
	}

	return maxLen
}
