package main

import (
	"reflect"
	"testing"
)

func TestLinerSearch02(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{1, 2, 2, 4, 5, 6, 7, 8, 9, 10}, 2, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, []int{}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, []int{}},
	}

	for _, tt := range tests {
		got := linerSearch02(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.result) {
			t.Errorf("got %v, want %v", got, tt.result)
		}
	}
}
