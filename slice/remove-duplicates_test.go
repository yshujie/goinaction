package slice_test

import (
	"testing"
)

// 通用断言：元素和顺序都必须一致
func assertSliceEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len mismatch: got=%d want=%d\n got=%v\nwant=%v", len(got), len(want), got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("element mismatch at %d: got=%v want=%v\n got=%v\nwant=%v", i, got[i], want[i], got, want)
		}
	}
}

func TestRemoveDuplicates_Int(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{42}, []int{42}},
		{"already_unique", []int{1, 2, 3}, []int{1, 2, 3}},
		{"duplicates_mixed", []int{1, 2, 1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{"all_same", []int{5, 5, 5, 5}, []int{5}},
		{"head_dups", []int{7, 7, 7, 8, 9}, []int{7, 8, 9}},
		{"tail_dups", []int{1, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"interleaved", []int{0, 1, 0, 1, 0, 1}, []int{0, 1}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveDuplicates(tc.in)
			assertSliceEqual(t, got, tc.want)
		})
	}
}

func TestRemoveDuplicates_Rune(t *testing.T) {
	tests := []struct {
		name string
		in   []rune
		want []rune
	}{
		{"ascii_letters", []rune("aabbccddeeff"), []rune("abcdef")},
		{"mix_letters_digits", []rune{'A', 'A', '1', '1', 'B'}, []rune{'A', '1', 'B'}},
		{"unicode_cn", []rune{'你', '你', '好', '好', '你'}, []rune{'你', '好'}},
		{"single_rune", []rune{'中'}, []rune{'中'}},
		{"empty", []rune{}, []rune{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveDuplicates(tc.in)
			assertSliceEqual(t, got, tc.want)
		})
	}
}

func TestRemoveDuplicates_String(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"empty", []string{}, []string{}},
		{"single", []string{"go"}, []string{"go"}},
		{"already_unique", []string{"a", "ab", "abc"}, []string{"a", "ab", "abc"}},
		{"duplicates_mixed", []string{"ab", "ab", "a", "b", "a", "ab"}, []string{"ab", "a", "b"}},
		{"all_same", []string{"x", "x", "x"}, []string{"x"}},
		{"head_tail_dups", []string{"k", "k", "m", "n", "n"}, []string{"k", "m", "n"}},
		{"unicode_strings", []string{"你", "你", "好", "世界", "好"}, []string{"你", "好", "世界"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveDuplicates(tc.in)
			assertSliceEqual(t, got, tc.want)
		})
	}
}
