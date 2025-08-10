package slice_test

import (
	"slices"
	"testing"
)

func TestRemoveSilceValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		val      int
		expected []int
	}{
		{
			name:     "普通情况：中间有多个目标值",
			input:    []int{1, 2, 3, 2, 4},
			val:      2,
			expected: []int{1, 3, 4},
		},
		{
			name:     "无目标值：返回原切片内容",
			input:    []int{1, 3, 4},
			val:      2,
			expected: []int{1, 3, 4},
		},
		{
			name:     "全是目标值：返回空切片",
			input:    []int{2, 2, 2},
			val:      2,
			expected: []int{},
		},
		{
			name:     "目标值在开头和结尾",
			input:    []int{2, 1, 3, 2},
			val:      2,
			expected: []int{1, 3},
		},
		{
			name:     "空切片",
			input:    []int{},
			val:      2,
			expected: []int{},
		},
		{
			name:     "单元素：是目标值",
			input:    []int{2},
			val:      2,
			expected: []int{},
		},
		{
			name:     "单元素：非目标值",
			input:    []int{1},
			val:      2,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建副本，防止测试之间数据污染
			inputCopy := append([]int(nil), tt.input...)
			result := RemoveSilceValue(inputCopy, tt.val)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
