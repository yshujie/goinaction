package base

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// StringToByteSlice 字符串->字节切片
func StringToByteSlice() {
	// 声明字符串
	var s string = "hello"

	// 将字符串强制转化为字节切片
	b := []byte(s)

	// 获取字节切片数据
	fmt.Print("b[0]: ", b[0])
	fmt.Println()
}

// ByteSlicsToString 字节切片->字符串
func ByteSlicsToString() {
	// 声明字节切片
	var b []byte = []byte{'h', 'e', 'l', 'l', 'o'}

	s := string(b)
	fmt.Print("string is ", s)
	fmt.Println()
}

// StringToRuneSlice 字符串->字符切片
func StringToRuneSlice() {
	// 声明字符串
	var s string = "hello 世界"

	// 字符串转字符切片
	r := []rune(s)

	fmt.Print("rune is ", r)
	fmt.Println()
}

// RuneSliceToString 字符切片->字符串
func RuneSliceToString() {
	// 声明字节切片
	var r []rune = []rune{'h', 'e', 'l', 'l', 'o', ' ', '世', '界'}

	// 字节切片强制转换为字符串
	s := string(r)

	// 输出字符串
	fmt.Print("string is ", s)
	fmt.Println()
}

// ForIterator for 循环遍历器
func ForIterator() {
	s := "Hello 世界"

	for i := 0; i <= len(s)-1; i++ {
		fmt.Printf("byte at %d: %x\n", i, s[i])
	}
}

// ForRangeIterator for range 循环遍历器
func ForRangeIterator() {
	s := "Hello 世界"

	for i, r := range s {
		fmt.Printf("rune at %d: %c (U+%04X)\n", i, r, r)
	}
}

// StringSplit 字符串切割
func StringSplit() {
	// 声明字符串
	s := "Go,python,Java"
	// 将字符串按照 "," 进行切分，返回字符串切片
	parts := strings.Split(s, ",")

	fmt.Print("parts: ", parts)
	fmt.Println()
}

// countStrLength 统计字符串长度
func CountStrLength(s string) {
	// 以字节维度，统计字符串长度
	fmt.Printf("s byte lenght by len(s): %d \n", len(s))
	fmt.Printf("s byte lenght by len([]byte(s)): %d \n", len([]byte(s)))

	// 以字符维度，统计字符串长度
	fmt.Printf("s charator lenght by utf8.RuneCountInString(s): %d \n", utf8.RuneCountInString(s))
	fmt.Printf("s charator lenght by len([]rune(s)): %d \n", len([]rune(s)))
}

// 遍历字符串的字符，并逐个输出
func StrIterator(s string) {
	for i, c := range s {
		// fmt.Printf("s: %s", string(c))
		fmt.Printf("index: %d; chart: %c; string: %s; type: %T; 0x: %x \n", i, c, string(c), c, c)
	}
}

// encodeRune 对 Rune 类型进行编码，Rune -> []byte
func EncodeRune(r rune) {
	fmt.Printf("the unicode charactor is %c \n", r)

	// 创建字节数组
	buf := make([]byte, 3)
	// 对 rune 进行 utf-8 编码
	utf8.EncodeRune(buf, r)

	// 十六进制（大写）
	fmt.Printf("utf-8 representation is 0x：%X \n ", buf)
}

// DecodeRune 对 byte 数组进行解码， []byte -> rune
func DecodeRune(buf []byte) {
	// 十六进制（大写）
	fmt.Printf("buf is 0x: %X \n", buf)

	r, s := utf8.DecodeRune(buf)
	fmt.Printf("rune chart is %c ; size is %d \n", r, s)
}

// 输出字符串的字节数组、字符数组
func PrintByteArrAndCharactorArr(s string) {
	s_bytes := []byte(s)
	fmt.Printf("string's byte arr is %X \n", s_bytes)
	// 以字节维度遍历字符串
	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d; byte: %X \n", i, s_bytes[i])
	}

	s_charactor := []rune(s)
	fmt.Printf("string's rune arr is %X, %c \n", s_charactor, s_charactor)
	for i, c := range s {
		fmt.Printf("index: %d; charator: %c; 16进制: %X \n", i, c, c)
	}
}

// 字符串拼接
func StrMerge(str ...string) string {
	if len(str) <= 3 {
		return strMergeLowSpeed(str...)
	} else {
		return strMergeHighSpeed(str...)
	}
}

// 低速方式合并字符串
func strMergeLowSpeed(str ...string) string {
	fmt.Println("use low speed merge function")
	result := ""
	for _, s := range str {
		result += s
	}

	return result
}

// 告诉方式合并字符串
func strMergeHighSpeed(str ...string) string {
	fmt.Println("use high speed merge function")
	// 字符串构造器
	var builder strings.Builder

	// 预分配 builder 缓冲区长度
	builder.Grow(10)

	for _, s := range str {
		builder.WriteString(s)
	}

	return builder.String()
}
