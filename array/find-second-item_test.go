package array_test

import (
	"testing"
)

func TestFindSecondItem(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expected    int
		expectError bool
	}{
		{
			name:        "正常情况 - 有第二大的元素",
			input:       []int{3, 1, 4, 5, 2},
			expected:    4,
			expectError: false,
		},
		{
			name:        "正常情况 - 所有元素相同",
			input:       []int{5, 5, 5, 5},
			expected:    5,
			expectError: false,
		},
		{
			name:        "正常情况 - 只有两个元素",
			input:       []int{10, 5},
			expected:    5,
			expectError: false,
		},
		{
			name:        "正常情况 - 负数元素",
			input:       []int{-3, -1, -4, -5, -2},
			expected:    -2,
			expectError: false,
		},
		{
			name:        "正常情况 - 混合正负数",
			input:       []int{-5, 10, -3, 8, 0},
			expected:    8,
			expectError: false,
		},
		{
			name:        "边界情况 - 空数组",
			input:       []int{},
			expected:    0,
			expectError: true,
		},
		{
			name:        "边界情况 - 只有一个元素",
			input:       []int{42},
			expected:    0,
			expectError: true,
		},
		{
			name:        "正常情况 - 递减序列",
			input:       []int{5, 4, 3, 2, 1},
			expected:    4,
			expectError: false,
		},
		{
			name:        "正常情况 - 递增序列",
			input:       []int{1, 2, 3, 4, 5},
			expected:    4,
			expectError: false,
		},
		{
			name:        "正常情况 - 包含重复元素",
			input:       []int{3, 3, 1, 4, 4, 5, 2},
			expected:    4,
			expectError: false,
		},
		{
			name:        "正常情况 - 负数",
			input:       []int{-5, -1, -3},
			expected:    -3,
			expectError: false,
		},
		{
			name:        "正常情况 - 包含多个最大值",
			input:       []int{5, 5, 3},
			expected:    3,
			expectError: false,
		},
		{
			name:        "特殊情况 - 全为负数",
			input:       []int{-10, -5, -15, -8},
			expected:    -8,
			expectError: false,
		},
		{
			name:        "特殊情况 - 包含零",
			input:       []int{0, 1, 0, 2, 0},
			expected:    1,
			expectError: false,
		},
		{
			name:        "特殊情况 - 大数组",
			input:       []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91},
			expected:    99,
			expectError: false,
		},
		{
			name:        "特殊情况 - 第二大的在开头",
			input:       []int{4, 5, 1, 2, 3},
			expected:    4,
			expectError: false,
		},
		{
			name:        "特殊情况 - 第二大的在结尾",
			input:       []int{1, 2, 3, 4, 5},
			expected:    4,
			expectError: false,
		},
		{
			name:        "特殊情况 - 所有元素都是第二大的",
			input:       []int{3, 3, 3, 3, 3},
			expected:    3,
			expectError: false,
		},
		{
			name:        "特殊情况 - 最大值在中间",
			input:       []int{1, 2, 5, 3, 4},
			expected:    4,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FindSecondItem(tt.input)

			// 检查错误情况
			if tt.expectError {
				if err == nil {
					t.Errorf("FindSecondItem(%v) 应该返回错误，但没有返回", tt.input)
				}
				return
			}

			// 检查正常情况
			if err != nil {
				t.Errorf("FindSecondItem(%v) 不应该返回错误，但返回了: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("FindSecondItem(%v) = %d, 期望 %d", tt.input, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkFindSecondItem(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSecondItem(input)
	}
}

// 测试不同大小数组的性能
func BenchmarkFindSecondItemDifferentSizes(b *testing.B) {
	testCases := []struct {
		name string
		size int
	}{
		{"小数组", 10},
		{"中等数组", 100},
		{"大数组", 1000},
		{"超大数组", 10000},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			input := make([]int, tc.size)
			for i := range input {
				input[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				FindSecondItem(input)
			}
		})
	}
}

// 测试特殊情况 - 所有元素相同
func BenchmarkFindSecondItemAllSame(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = 42
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSecondItem(input)
	}
}

// 测试特殊情况 - 递减序列
func BenchmarkFindSecondItemDecreasing(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = 1000 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSecondItem(input)
	}
}
