package array_test

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "正常情况 - 偶数长度数组",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name:     "正常情况 - 奇数长度数组",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "边界情况 - 空数组",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "边界情况 - 单个元素",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "边界情况 - 两个元素",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "特殊情况 - 包含重复元素",
			input:    []int{1, 2, 2, 1},
			expected: []int{1, 2, 2, 1},
		},
		{
			name:     "特殊情况 - 包含负数",
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{3, 2, 1, 0, -1, -2, -3},
		},
		{
			name:     "特殊情况 - 包含零",
			input:    []int{0, 1, 0, 2, 0},
			expected: []int{0, 2, 0, 1, 0},
		},
		{
			name:     "大数组测试",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入切片的副本，因为 Reverse 函数会修改原切片
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 调用 Reverse 函数
			Reverse(input)

			// 检查结果
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("Reverse() = %v, 期望 %v", input, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkReverse(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 创建输入切片的副本
		testInput := make([]int, len(input))
		copy(testInput, input)
		Reverse(testInput)
	}
}
