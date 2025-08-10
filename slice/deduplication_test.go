package slice_test

import (
	"reflect"
	"testing"
)

func TestDeduplication(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "No duplicates",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "All same",
			input: []string{"a", "a", "a"},
			want:  []string{"a"},
		},
		{
			name:  "Adjacent duplicates",
			input: []string{"a", "a", "b", "b", "c", "a", "a"},
			want:  []string{"a", "b", "c", "a"},
		},
		{
			name:  "Adjacent Chinese",
			input: []string{"你", "你", "们", "好", "呀", "吗", "呀", "呀"},
			want:  []string{"你", "们", "好", "呀", "吗", "呀"},
		},
		{
			name:  "Single element",
			input: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "Empty slice",
			input: []string{},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Deduplication(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deduplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
