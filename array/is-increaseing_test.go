package array_test

import (
	"testing"
)

func TestIsIncreasing(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"Strictly increasing", []int{1, 2, 3, 4, 5}, true},
		{"Non-decreasing with equal elements", []int{1, 2, 2, 3, 4}, true},
		{"Decreasing array", []int{5, 4, 3, 2, 1}, false},
		{"Single element", []int{1}, true},
		{"Empty array", []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsIncreasing(tt.arr)
			if got != tt.want {
				t.Errorf("IsIncreasing(%v) = %v; want %v", tt.arr, got, tt.want)
			}
		})
	}
}
