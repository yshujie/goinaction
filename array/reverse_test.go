package array_test

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input [5]int
		want  [5]int
	}{
		{
			name:  "positive order",
			input: [5]int{1, 2, 3, 4, 5},
			want:  [5]int{5, 4, 3, 2, 1},
		},
		{
			name:  "reverse order",
			input: [5]int{5, 4, 3, 2, 1},
			want:  [5]int{1, 2, 3, 4, 5},
		},
		{
			name:  "disorder",
			input: [5]int{9, 4, 10, 86, 10},
			want:  [5]int{10, 86, 10, 4, 9},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the original
			got := tc.input
			Reverse(&got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Reverse(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
