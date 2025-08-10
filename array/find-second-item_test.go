package array_test

import (
	"testing"
)

func TestFindSecondItem(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "正常情况 - 有第二大的元素",
			input:    []int{3, 1, 4, 5, 2},
			expected: 4,
		},
		{
			name:     "正常情况 - 所有元素相同",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "正常情况 - 只有两个元素",
			input:    []int{10, 5},
			expected: 5,
		},
		{
			name:     "正常情况 - 负数元素",
			input:    []int{-3, -1, -4, -5, -2},
			expected: -2,
		},
		{
			name:     "正常情况 - 混合正负数",
			input:    []int{-5, 10, -3, 8, 0},
			expected: 8,
		},
		{
			name:     "边界情况 - 空数组",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "边界情况 - 只有一个元素",
			input:    []int{42},
			expected: 0,
		},
		{
			name:     "正常情况 - 递减序列",
			input:    []int{5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "正常情况 - 递增序列",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "正常情况 - 包含重复元素",
			input:    []int{3, 3, 1, 4, 4, 5, 2},
			expected: 4,
		},
		{
			name:     "正常情况 - 负数",
			input:    []int{-5, -1, -3},
			expected: -3,
		},
		{
			name:     "正常情况 - 包含多个最大值",
			input:    []int{5, 5, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindSecondItem(tt.input)
			if result != tt.expected {
				t.Errorf("FindSecondItem(%v) = %d, 期望 %d", tt.input, result, tt.expected)
			}
		})
	}
}
