package base

import "fmt"

/**
✅ 题目五：找出切片中最长的连续递增子序列长度
输入: []int{1, 2, 2, 3, 4, 1, 2, 3, 4, 5}
输出: 5   // 对应子序列 [1 2 3 4 5]

✍ 要求
- 返回 int 类型结果
- 时间复杂度 O(n)
*/

/*
*
方案一：从前向后遍历，从第一个元素开始排查，检查 back-prev == 1 ？直至 back-prev ！= 1，此时当前子串长度为 back index - prev index + 1
出现后，记录当前结果，截取 [back index:]，重新开始，继续向后查询
*/
func CountMaxSubSliceCnt1(s []int) int {
	fmt.Println("in CountMaxSubSlice, s: ", s)
	count := 1
	for i := 0; i <= len(s)-2; i++ {
		if s[i+1]-s[i] == 1 {
			count++
		} else {
			// 截取新的 slice
			leftSlice := s[i+1:]

			// 统计后半段 slice 的最长子穿长度
			backCount := CountMaxSubSliceCnt1(leftSlice)
			if backCount > count {
				count = backCount
			}

			// 停止当前子串的循环计数
			break
		}
	}

	return count
}

/*
*
方案二：使用两个变量，分别记录 子串长度 和 当前子串长度
循环切片，检查 back - prev == 1, 若是，则当前子串长度+1；若非，则重新计数。
重新计数时，判断 当前子串长度 > 子串长度 ？ 是，则用当前子串长度覆盖子串长度，否则遗弃
*/
func CountMaxSubSliceCnt2(s []int) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxCnt := 1
	currCnt := 1
	for i := 1; i <= len(s)-1; i++ {
		if s[i]-s[i-1] == 1 {
			currCnt++
		} else {
			currCnt = 1
		}

		if currCnt > maxCnt {
			maxCnt = currCnt
		}
	}

	return maxCnt
}
