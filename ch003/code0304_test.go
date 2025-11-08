package main

import (
	"math"
	"testing"
)

func TestSearchMinValidPairSum(t *testing.T) {

	tests := []struct {
		nums1    []int
		nums2    []int
		k        int
		result   int
		is_error bool
	}{
		// 通常ケース（ランダムな並び）
		{[]int{7, 2, 9, 1, 5}, []int{3, 8, 4, 1, 6}, 6, 6, false},
		{[]int{15, 3, 22, 8, 11}, []int{7, 19, 2, 13}, 12, 13, false},
		{[]int{5, 18, 2, 9}, []int{12, 3, 20, 7, 15}, 10, 12, false},

		// 負の数を含む（バラバラ）
		{[]int{-3, 8, -12, 5, 0}, []int{4, -7, 11, -2}, 0, 1, false},
		{[]int{-15, 3, -8, 20}, []int{-5, 12, -18, 7}, -10, -8, false},

		// エッジケース
		{[]int{42}, []int{17}, 50, 59, false},
		{[]int{3, 9, 3, 6}, []int{5, 5, 2, 5}, 8, 8, false},
		{[]int{100, 25, 73, 41}, []int{88, 12, 55}, 50, 53, false},

		// エラーケース
		{[]int{}, []int{5, 12, 8}, 10, math.MaxInt, true},
		{[]int{7, 3, 9}, []int{}, 5, math.MaxInt, true},
		{[]int{}, []int{}, 10, math.MaxInt, true},
		{[]int{1, 3, 2}, []int{4, 2, 5}, 100, math.MaxInt, true},
	}

	for _, tt := range tests {
		sum_value, is_error := searchMinValidPairSum(tt.nums1, tt.nums2, tt.k)

		if is_error != tt.is_error {
			t.Errorf("is_error : got %v, want %v", is_error, tt.is_error)
		}

		if sum_value != tt.result {
			t.Errorf("min_value : got %v, want %v", sum_value, tt.result)
		}
	}
}
