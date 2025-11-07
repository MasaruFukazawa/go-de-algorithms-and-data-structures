package main

import "testing"

func TestLinerSearch(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, false},
	}

	for _, tt := range tests {
		got := linerSearch(tt.nums, tt.target)
		if tt.result != got {
			t.Errorf("got %v, want %v", got, tt.result)
		}
	}
}
