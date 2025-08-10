package map_test

import (
	"slices"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "两个数组的交集 - 基本测试",
			input:    [][]int{{1, 2, 2, 1}, {2, 2}},
			expected: []int{2},
		},
		{
			name:     "两个数组的交集 - 多个共同元素",
			input:    [][]int{{1, 2, 3, 4}, {2, 3, 5, 6}},
			expected: []int{2, 3},
		},
		{
			name:     "两个数组的交集 - 无交集",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: []int{},
		},
		{
			name:     "两个数组的交集 - 完全重叠",
			input:    [][]int{{1, 2, 3}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "两个数组的交集 - 空数组",
			input:    [][]int{{}, {1, 2, 3}},
			expected: []int{},
		},
		{
			name:     "两个数组的交集 - 两个空数组",
			input:    [][]int{{}, {}},
			expected: []int{},
		},
		{
			name:     "两个数组的交集 - 重复元素",
			input:    [][]int{{1, 1, 1, 2, 2}, {1, 2, 2, 3}},
			expected: []int{1, 2},
		},
		{
			name:     "两个数组的交集 - 负数",
			input:    [][]int{{-1, -2, 1, 2}, {-1, 1, 3}},
			expected: []int{-1, 1},
		},
		{
			name:     "两个数组的交集 - 零值",
			input:    [][]int{{0, 1, 2}, {0, 2, 3}},
			expected: []int{0, 2},
		},
		{
			name:     "三个数组的交集",
			input:    [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}},
			expected: []int{3, 4},
		},
		{
			name:     "三个数组的交集 - 无交集",
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{},
		},
		{
			name:     "三个数组的交集 - 完全重叠",
			input:    [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "四个数组的交集",
			input:    [][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5, 6}, {3, 4, 5, 6, 7}, {4, 5, 6, 7, 8}},
			expected: []int{4, 5},
		},
		{
			name:     "单个数组",
			input:    [][]int{{1, 2, 3, 4, 5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "零个数组",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "大数字测试",
			input:    [][]int{{1000, 2000, 3000}, {2000, 3000, 4000}},
			expected: []int{2000, 3000},
		},
		{
			name:     "包含空数组的多个数组",
			input:    [][]int{{1, 2, 3}, {}, {2, 3, 4}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.input...)

			// 对结果进行排序以便比较（因为顺序不重要）
			sort.Ints(result)
			sort.Ints(tt.expected)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("Intersection() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkIntersection(b *testing.B) {
	input := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
	}

	for i := 0; i < b.N; i++ {
		Intersection(input...)
	}
}

// 测试大输入的性能
func TestIntersectionLargeInput(t *testing.T) {
	// 创建三个大数组
	arr1 := make([]int, 1000)
	arr2 := make([]int, 1000)
	arr3 := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		arr1[i] = i
		arr2[i] = i + 500 // 与arr1有500个重叠元素
		arr3[i] = i + 750 // 与arr1有250个重叠元素，与arr2有250个重叠元素
	}

	result := Intersection(arr1, arr2, arr3)

	// 期望的交集应该是 [750, 751, ..., 999]
	expected := make([]int, 250)
	for i := 0; i < 250; i++ {
		expected[i] = 750 + i
	}

	sort.Ints(result)
	sort.Ints(expected)

	if !slices.Equal(result, expected) {
		t.Errorf("大输入测试失败: 得到 %v, 期望 %v", result, expected)
	}
}

// 测试边界情况
func TestIntersectionEdgeCases(t *testing.T) {
	// 测试包含重复元素的情况
	result := Intersection([]int{1, 1, 1, 2, 2, 3}, []int{1, 2, 2, 2, 3, 3})
	sort.Ints(result)
	expected := []int{1, 2, 3}
	sort.Ints(expected)

	if !slices.Equal(result, expected) {
		t.Errorf("重复元素测试失败: 得到 %v, 期望 %v", result, expected)
	}

	// 测试只有一个元素的数组
	result = Intersection([]int{42}, []int{42, 43, 44})
	expected = []int{42}

	if !slices.Equal(result, expected) {
		t.Errorf("单元素测试失败: 得到 %v, 期望 %v", result, expected)
	}
}
