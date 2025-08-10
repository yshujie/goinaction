package map_test

/**
✅ 题目：判断两个字符串是否为字母异位词
输入两个字符串 s1 和 s2，判断它们是否为字母异位词（字符相同但顺序不同）。

示例：
输入："listen" 和 "silent" → 输出：true

考点：
	•	map[rune]int 计数法
	•	遍历字符 + 增减计数
	•	Unicode 兼容
*/

/**
思路：
- 字母异位词的定义：1.字符长度相同；2.字符出现的频率相同；3.字符出现的顺序不同；
- 判断方式：
	1. 若字符串相等，则不符合定义 3 ，为 false；
	2. 若字符串长度不同，则不符合定义 1，为 false；
	3. 若满足字符串不相等、字符串长度相同、字符出现的频率相同，则说明是字母异位词
*/

func isAnagram(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}
	if len([]rune(s1)) != len([]rune(s2)) {
		return false
	}

	// 字符统计表，s1 累加，s2 累减
	freq := make(map[rune]int)
	for _, v := range s1 {
		freq[v]++
	}
	for _, v := range s2 {
		freq[v]--

		// 若该字符的次数小于 0，则说明 s1 中未累加，则直接判定“非字母异位词”
		if freq[v] < 0 {
			return false
		}
	}

	return true
}
