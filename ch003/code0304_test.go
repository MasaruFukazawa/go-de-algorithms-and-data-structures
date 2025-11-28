package main

import (
	"math"
	"testing"
)

func TestSearchMinValidPairSum(t *testing.T) {

	tests := []struct {
		name    string
		nums1   []int
		nums2   []int
		k       int
		want    int
		wantErr bool
	}{
		// 通常ケース（ランダムな並び）
		{"Normal case 1", []int{7, 2, 9, 1, 5}, []int{3, 8, 4, 1, 6}, 6, 6, false},
		{"Normal case 2", []int{15, 3, 22, 8, 11}, []int{7, 19, 2, 13}, 12, 13, false},
		{"Normal case 3", []int{5, 18, 2, 9}, []int{12, 3, 20, 7, 15}, 10, 12, false},

		// 負の数を含む（バラバラ）
		{"Negative numbers 1", []int{-3, 8, -12, 5, 0}, []int{4, -7, 11, -2}, 0, 1, false},
		{"Negative numbers 2", []int{-15, 3, -8, 20}, []int{-5, 12, -18, 7}, -10, -8, false},

		// エッジケース
		{"Single elements", []int{42}, []int{17}, 50, 59, false},
		{"Duplicates", []int{3, 9, 3, 6}, []int{5, 5, 2, 5}, 8, 8, false},
		{"Large numbers", []int{100, 25, 73, 41}, []int{88, 12, 55}, 50, 53, false},

		// エラーケース
		{"Empty nums1", []int{}, []int{5, 12, 8}, 10, math.MaxInt, true},
		{"Empty nums2", []int{7, 3, 9}, []int{}, 5, math.MaxInt, true},
		{"Both empty", []int{}, []int{}, 10, math.MaxInt, true},
		{"No valid pair", []int{1, 3, 2}, []int{4, 2, 5}, 100, math.MaxInt, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := searchMinValidPairSum(tt.nums1, tt.nums2, tt.k)

			if err != tt.wantErr {
				t.Errorf("searchMinValidPairSum() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("searchMinValidPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
