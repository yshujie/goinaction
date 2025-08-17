package array_test

import (
	"reflect"
	"testing"
)

func TestRotatingLeft(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "正常情况 - 向左旋转3位",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{4, 5, 6, 7, 1, 2, 3},
		},
		{
			name:     "正常情况 - 向左旋转1位",
			input:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{2, 3, 4, 5, 1},
		},
		{
			name:     "正常情况 - 向左旋转2位",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "边界情况 - 旋转0位",
			input:    []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "边界情况 - 旋转数组长度",
			input:    []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "边界情况 - 旋转超过数组长度",
			input:    []int{1, 2, 3, 4, 5},
			k:        7,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "边界情况 - 空数组",
			input:    []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "边界情况 - 单个元素",
			input:    []int{42},
			k:        1,
			expected: []int{42},
		},
		{
			name:     "边界情况 - 两个元素",
			input:    []int{1, 2},
			k:        1,
			expected: []int{2, 1},
		},
		{
			name:     "特殊情况 - 包含负数",
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			k:        2,
			expected: []int{-1, 0, 1, 2, 3, -3, -2},
		},
		{
			name:     "特殊情况 - 包含重复元素",
			input:    []int{1, 1, 2, 2, 3, 3},
			k:        2,
			expected: []int{2, 2, 3, 3, 1, 1},
		},
		{
			name:     "特殊情况 - 包含零",
			input:    []int{0, 1, 0, 2, 0},
			k:        3,
			expected: []int{2, 0, 0, 1, 0},
		},
		{
			name:     "大数组测试 - 旋转一半",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        5,
			expected: []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
		},
		{
			name:     "负数旋转 - 应该等同于正数旋转",
			input:    []int{1, 2, 3, 4, 5},
			k:        -2,
			expected: []int{4, 5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入切片的副本，因为 RotatingLeft 函数会修改原切片
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			// 调用 RotatingLeft 函数
			result := RotatingLeft(input, tt.k)

			// 检查结果
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RotatingLeft() = %v, 期望 %v", result, tt.expected)
			}

			// 检查原切片是否被正确修改
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("原切片未被正确修改，实际 = %v, 期望 %v", input, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkRotatingLeft(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 创建输入切片的副本
		testInput := make([]int, len(input))
		copy(testInput, input)
		RotatingLeft(testInput, 333)
	}
}

// 测试不同旋转步数的性能
func BenchmarkRotatingLeftDifferentK(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	testCases := []struct {
		name string
		k    int
	}{
		{"旋转1位", 1},
		{"旋转10位", 10},
		{"旋转100位", 100},
		{"旋转500位", 500},
		{"旋转999位", 999},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testInput := make([]int, len(input))
				copy(testInput, input)
				RotatingLeft(testInput, tc.k)
			}
		})
	}
}
