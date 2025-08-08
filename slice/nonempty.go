package slice_test

/**
✅ 题目：去除字符串切片中的空字符
实现一个函数 nonempty，它接收一个 []string 切片作为输入，返回一个新的 []string 切片，其中只包含非空字符串，且保持原有顺序不变。

要求：
	1.	不能修改原切片的内容（返回新的切片）。
	2.	切片中元素顺序保持不变。
	3.	只保留那些不等于 "" 的字符串。

函数签名：
func nonempty(strings []string) []string

输入：
[]string{"one", "", "three"}

输出：
[]string{"one", "three"}

考点：
	•	切片遍历
	•	条件过滤
	•	新切片构造
	•	保持输入不变的副本处理
*/

/*
*
思路：新建切片法
  - 创建一个空切片
  - 将原切片从前向后遍历，检查每一个元素
  - 若元素不等于 ""，则 append 入新 slice
  - 若元素等于 ""，则跳过
*/
func Nonempty(strings []string) []string {
	result := []string{}

	for _, v := range strings {
		if v == "" {
			continue
		}

		result = append(result, v)
	}

	return result
}

/*
*
思路：元素搬运法
  - 这种思路，其实有 “读” 和 “写” 两个进度
  - 声明变量 i := 0，是“写”进度
  - 从左向右遍历 strings，是 “读” 进度
  - 读的时候，发现元素不为 "" 则将元素写入原切片，并更新写进度；若元素为 "" 则跳过，也不更新写进度
  - 最终返回 strings[:i]
*/
func Nonempty2(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}

	return strings[:i]
}
