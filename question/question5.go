package question

import "strings"

/**
题目：字符串替换问题
请编写一个方法，将字符串中的空格全部替换为“%20”。

要求：
- 假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000);
- 给定一个string为原始的串，返回替换后的string。
*/

func ReplaceBlank(str string) string {
	var strBuilder strings.Builder

	for _, v := range str {
		if v == ' ' {
			strBuilder.WriteString("%20")
		} else {
			strBuilder.WriteRune(v)
		}
	}

	return strBuilder.String()
}
