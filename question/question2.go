package question

import "strings"

/**
题目：判断字符串中字符是否全都不同
请实现一个算法，确定一个字符串的所有字符是否全都不同。

要求：
1. 不允许使用额外的存储结构
2. 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
3. 保证字符串中的字符为 ASCII 字符
4. 字符串的长度小于等于 3000
*/

func IsUniqueString(s string) bool {
	if len(s) > 3000 { // 字符串长度超过限制
		return false
	}

	for i, v := range s {
		if v > 127 { // 出现非 ASCII 字符
			return false
		}

		// 字符索引值 与 当前迭代器所有不同，说明出现多次
		if strings.Index(s, string(v)) != i {
			return false
		}
	}

	return true
}
