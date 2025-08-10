package map_test

import (
	"slices"
	"sort"
	"testing"
)

func TestFindMostFrequentItem(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantCount int
		wantItems []int
	}{
		{
			name:      "空切片",
			input:     []int{},
			wantCount: 0,
			wantItems: []int{},
		},
		{
			name:      "单个元素",
			input:     []int{1},
			wantCount: 1,
			wantItems: []int{1},
		},
		{
			name:      "两个不同元素",
			input:     []int{1, 2},
			wantCount: 1,
			wantItems: []int{1, 2},
		},
		{
			name:      "两个相同元素",
			input:     []int{1, 1},
			wantCount: 2,
			wantItems: []int{1},
		},
		{
			name:      "多个元素，一个出现最多",
			input:     []int{1, 2, 3, 2, 4, 2, 5},
			wantCount: 3,
			wantItems: []int{2},
		},
		{
			name:      "多个元素，多个并列最多",
			input:     []int{1, 3, 2, 3, 4, 3, 2, 2},
			wantCount: 3,
			wantItems: []int{2, 3},
		},
		{
			name:      "所有元素出现次数相同",
			input:     []int{1, 2, 3, 4},
			wantCount: 1,
			wantItems: []int{1, 2, 3, 4},
		},
		{
			name:      "负数元素",
			input:     []int{-1, -2, -1, -3, -2, -1},
			wantCount: 3,
			wantItems: []int{-1},
		},
		{
			name:      "零值元素",
			input:     []int{0, 1, 0, 2, 0},
			wantCount: 3,
			wantItems: []int{0},
		},
		{
			name:      "大数字",
			input:     []int{1000, 1000, 999, 1000, 999},
			wantCount: 3,
			wantItems: []int{1000},
		},
		{
			name:      "复杂情况：多个并列最多",
			input:     []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
			wantCount: 2,
			wantItems: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotItems := FindMostFrequentItem(tt.input)

			// 检查次数
			if gotCount != tt.wantCount {
				t.Errorf("FindMostFrequentItem() count = %v, want %v", gotCount, tt.wantCount)
			}

			// 对结果进行排序以便比较（因为顺序不重要）
			sort.Ints(gotItems)
			sort.Ints(tt.wantItems)

			// 检查元素列表
			if !slices.Equal(gotItems, tt.wantItems) {
				t.Errorf("FindMostFrequentItem() items = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}

// 基准测试
func BenchmarkFindMostFrequentItem(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		FindMostFrequentItem(input)
	}
}

// 测试大输入的性能
func TestFindMostFrequentItemLargeInput(t *testing.T) {
	// 创建一个包含1000个元素的大切片
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i % 100 // 0-99的数字，每个数字出现10次
	}

	// 添加一些出现次数更多的数字
	for i := 0; i < 50; i++ {
		input = append(input, 999) // 数字999出现50次
	}

	count, items := FindMostFrequentItem(input)

	if count != 50 {
		t.Errorf("期望最大出现次数为50，实际得到 %d", count)
	}

	if len(items) != 1 || items[0] != 999 {
		t.Errorf("期望出现最多的元素为999，实际得到 %v", items)
	}
}
