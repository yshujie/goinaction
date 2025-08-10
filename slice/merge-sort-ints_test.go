package slice_test

import (
	"reflect"
	"testing"
)

func TestMergeSortedInts(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "正常情况 - 两个非空切片",
			a:        []int{1, 3, 5},
			b:        []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "正常情况 - 有重复元素",
			a:        []int{1, 2, 3},
			b:        []int{2, 3, 4},
			expected: []int{1, 2, 2, 3, 3, 4},
		},
		{
			name:     "边界情况 - a 为空切片",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "边界情况 - b 为空切片",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "边界情况 - 两个都为空切片",
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "特殊情况 - a 长度大于 b",
			a:        []int{1, 3, 5, 7, 9},
			b:        []int{2, 4},
			expected: []int{1, 2, 3, 4, 5, 7, 9},
		},
		{
			name:     "特殊情况 - b 长度大于 a",
			a:        []int{1, 3},
			b:        []int{2, 4, 6, 8, 10},
			expected: []int{1, 2, 3, 4, 6, 8, 10},
		},
		{
			name:     "特殊情况 - 负数元素",
			a:        []int{-3, -1, 1},
			b:        []int{-2, 0, 2},
			expected: []int{-3, -2, -1, 0, 1, 2},
		},
		{
			name:     "特殊情况 - 相同元素",
			a:        []int{1, 1, 1},
			b:        []int{1, 1, 1},
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "特殊情况 - 单元素切片",
			a:        []int{5},
			b:        []int{3},
			expected: []int{3, 5},
		},
		{
			name:     "特殊情况 - 一个单元素，一个多元素",
			a:        []int{5},
			b:        []int{1, 3, 7},
			expected: []int{1, 3, 5, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSortedInts(tt.a, tt.b)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSortedInts(%v, %v) = %v, 期望 %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// 基准测试
func BenchmarkMergeSortedInts(b *testing.B) {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15}
	slice := []int{2, 4, 6, 8, 10, 12, 14, 16}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortedInts(a, slice)
	}
}

// 基准测试 - 大切片
func BenchmarkMergeSortedIntsLarge(b *testing.B) {
	// 创建两个大切片
	a := make([]int, 1000)
	b_slice := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		a[i] = i * 2         // 0, 2, 4, 6, ...
		b_slice[i] = i*2 + 1 // 1, 3, 5, 7, ...
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortedInts(a, b_slice)
	}
}
