package question

/**
题目：判断两个给定的字符串排序后是否一致
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

要求：
- 规定【大小写为不同字符】；
- 给定一个string s1 和 一个string s2，请返回一个bool，代表两串是否重新排列后可相同；
- 保证两串的长度都小于等于5000。
*/

func IsRegroupStr(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) > 5000 || len(s2) > 5000 {
		return false
	}

	m := make(map[rune]int)
	for _, letter := range s1 {
		m[letter] += 1
	}
	for _, letter := range s2 {
		m[letter] -= 1
	}
	for _, num := range m {
		if num != 0 {
			return false
		}
	}

	return true
}
