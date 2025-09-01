package question

/**
题目：翻转字符串
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

要求：
1. 给定一个string，请返回一个string，为翻转后的字符串。
2. 保证字符串的长度小于等于5000。
*/

func ReverString(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}

	return string(r)
}
