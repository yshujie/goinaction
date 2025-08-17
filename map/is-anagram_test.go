package map_test

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "基本字母异位词测试",
			s1:       "listen",
			s2:       "silent",
			expected: true,
		},
		{
			name:     "另一个基本测试",
			s1:       "anagram",
			s2:       "nagaram",
			expected: true,
		},
		{
			name:     "相同字符串 - 应该返回false",
			s1:       "hello",
			s2:       "hello",
			expected: false,
		},
		{
			name:     "长度不同 - 应该返回false",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "长度相同但字符不同",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "空字符串",
			s1:       "",
			s2:       "",
			expected: false,
		},
		{
			name:     "一个空字符串",
			s1:       "hello",
			s2:       "",
			expected: false,
		},
		{
			name:     "另一个空字符串",
			s1:       "",
			s2:       "hello",
			expected: false,
		},
		{
			name:     "单个字符相同",
			s1:       "a",
			s2:       "a",
			expected: false,
		},
		{
			name:     "单个字符不同",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "包含重复字符的异位词",
			s1:       "aacc",
			s2:       "ccaa",
			expected: true,
		},
		{
			name:     "包含重复字符但不是异位词",
			s1:       "aacc",
			s2:       "ccab",
			expected: false,
		},
		{
			name:     "包含数字的异位词",
			s1:       "12345",
			s2:       "54321",
			expected: true,
		},
		{
			name:     "包含特殊字符的异位词",
			s1:       "!@#$%",
			s2:       "%$#@!",
			expected: true,
		},
		{
			name:     "包含空格和标点的异位词",
			s1:       "hello world!",
			s2:       "world! hello",
			expected: true,
		},
		{
			name:     "包含大写字母的异位词",
			s1:       "Hello",
			s2:       "oHell",
			expected: true,
		},
		{
			name:     "大小写混合的异位词",
			s1:       "Hello",
			s2:       "hEllo",
			expected: false,
		},
		{
			name:     "包含Unicode字符的异位词",
			s1:       "café",
			s2:       "éfac",
			expected: true,
		},
		{
			name:     "包含中文字符的异位词",
			s1:       "你好世界",
			s2:       "世界你好",
			expected: true,
		},
		{
			name:     "包含emoji的异位词",
			s1:       "😀😃😄",
			s2:       "😄😀😃",
			expected: true,
		},
		{
			name:     "长字符串的异位词",
			s1:       "abcdefghijklmnopqrstuvwxyz",
			s2:       "zyxwvutsrqponmlkjihgfedcba",
			expected: true,
		},
		{
			name:     "包含重复字符的长字符串",
			s1:       "aabbccddee",
			s2:       "eeddccbbaa",
			expected: true,
		},
		{
			name:     "字符频率不同",
			s1:       "aab",
			s2:       "abb",
			expected: false,
		},
		{
			name:     "字符频率相同但字符不同",
			s1:       "aab",
			s2:       "ccd",
			expected: false,
		},
		{
			name:     "包含制表符和换行符",
			s1:       "hello\tworld\n",
			s2:       "world\nhello\t",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAnagram(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkIsAnagram(b *testing.B) {
	s1 := "abcdefghijklmnopqrstuvwxyz"
	s2 := "zyxwvutsrqponmlkjihgfedcba"

	for i := 0; i < b.N; i++ {
		IsAnagram(s1, s2)
	}
}

// 测试大输入的性能
func TestIsAnagramLargeInput(t *testing.T) {
	// 创建两个大字符串
	s1 := "abcdefghijklmnopqrstuvwxyz"
	s2 := "zyxwvutsrqponmlkjihgfedcba"

	// 重复这个模式来创建大字符串
	for i := 0; i < 38; i++ { // 38 * 26 = 988 字符
		s1 += "abcdefghijklmnopqrstuvwxyz"
		s2 += "zyxwvutsrqponmlkjihgfedcba"
	}

	// 这两个字符串应该是异位词
	result := IsAnagram(s1, s2)
	if !result {
		t.Errorf("大输入测试失败: 期望 true，得到 %v", result)
	}

	// 测试不是异位词的情况
	s2 = s1[:len(s1)-1] + "x" // 修改最后一个字符
	result = IsAnagram(s1, s2)
	if result {
		t.Errorf("大输入测试失败: 期望 false，得到 %v", result)
	}
}

// 测试边界情况
func TestIsAnagramEdgeCases(t *testing.T) {
	// 测试非常长的字符串
	longStr1 := "abcdefghijklmnopqrstuvwxyz"
	longStr2 := "zyxwvutsrqponmlkjihgfedcba"

	// 重复这个模式来创建长字符串
	for i := 0; i < 384; i++ { // 384 * 26 = 9984 字符
		longStr1 += "abcdefghijklmnopqrstuvwxyz"
		longStr2 += "zyxwvutsrqponmlkjihgfedcba"
	}

	result := IsAnagram(longStr1, longStr2)
	if !result {
		t.Errorf("长字符串测试失败: 期望 true，得到 %v", result)
	}

	// 测试包含大量重复字符的字符串
	repeatStr1 := ""
	repeatStr2 := ""

	for i := 0; i < 1000; i++ {
		repeatStr1 += "a"
		repeatStr2 += "a"
	}

	result = IsAnagram(repeatStr1, repeatStr2)
	if result {
		t.Errorf("重复字符测试失败: 期望 false（相同字符串），得到 %v", result)
	}
}

// 测试对称性
func TestIsAnagramSymmetry(t *testing.T) {
	testCases := []struct {
		s1 string
		s2 string
	}{
		{"listen", "silent"},
		{"anagram", "nagaram"},
		{"hello world!", "world! hello"},
		{"你好世界", "世界你好"},
	}

	for _, tc := range testCases {
		// 测试 s1 和 s2
		result1 := IsAnagram(tc.s1, tc.s2)
		// 测试 s2 和 s1（应该得到相同结果）
		result2 := IsAnagram(tc.s2, tc.s1)

		if result1 != result2 {
			t.Errorf("对称性测试失败: IsAnagram(%q, %q) = %v, IsAnagram(%q, %q) = %v",
				tc.s1, tc.s2, result1, tc.s2, tc.s1, result2)
		}
	}
}
