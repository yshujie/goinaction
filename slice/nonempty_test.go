package slice_test

import (
	"reflect"
	"testing"
)

func TestNonempty(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"mixed empty", []string{"one", "", "three"}, []string{"one", "three"}},
		{"all empty", []string{"", "", ""}, []string{}},
		{"nil slice", nil, []string{}},
		{"no empty", []string{"a", "b"}, []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Nonempty(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nonempty(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNonempty2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"mixed empty", []string{"one", "", "three"}, []string{"one", "three"}},
		{"all empty", []string{"", "", ""}, []string{}},
		{"nil slice", nil, []string{}},
		{"no empty", []string{"a", "b"}, []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)
			got := Nonempty2(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nonempty2(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
