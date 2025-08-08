package string_test

/*
*
✅ 题目：反转字符串
编写函数 ReverseString(s string) string，使用 for 循环将字符串反转并返回。

输入示例：
fmt.Println(ReverseString("golang"))

输出：
gnalog
*/
func ReverseString(s string) string {
	result := ""

	for _, char := range s {
		result = string(char) + result
	}

	return result
}

// ReverseString2 反转字符串
// 使用 rune 切片的方式
func ReverseString2(s string) string {
	// 将字符串转为字符切片
	s_rune := []rune(s)
	lenght := len(s_rune)

	for i := 0; i < lenght/2; i++ {
		s_rune[i], s_rune[lenght-1-i] = s_rune[lenght-1-i], s_rune[i]
	}

	return string(s_rune)
}
