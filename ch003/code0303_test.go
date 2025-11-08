package main

import (
	"math"
	"testing"
)

func TestSearchMinValue(t *testing.T) {

	tests := []struct {
		nums     []int
		result   int
		is_error bool
	}{
		{[]int{1, 2, 2, 4, 5, 6, 7, 8, 9, 10}, 1, false}, // 先頭が最小
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1, false}, // 末尾が最小
		{[]int{5, 3, 8, 1, 9, 2, 7}, 1, false},           // 中間が最小
		{[]int{-5, -1, 0, 3, 7}, -5, false},              // 負の数を含む
		{[]int{100}, 100, false},                         // 要素が1つ
		{[]int{5, 5, 5, 5, 5}, 5, false},                 // 全て同じ値
		{[]int{-10, -20, -5, -15}, -20, false},           // 全て負の数
		{[]int{}, math.MaxInt, true},                     // 空配列
	}

	for _, tt := range tests {
		min_value, is_error := searchMinValue(tt.nums)

		if is_error != tt.is_error {
			t.Errorf("is_error : got %v, want %v", is_error, tt.is_error)
		}

		if min_value != tt.result {
			t.Errorf("min_value : got %v, want %v", min_value, tt.result)
		}
	}
}
