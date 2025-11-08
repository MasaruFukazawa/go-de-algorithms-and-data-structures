package main

import "math"

func searchMinValue(nums []int) (int, bool) {

	if len(nums) == 0 {
		return math.MaxInt, true
	}

	result := math.MaxInt

	for _, num := range nums {
		if result > num {
			result = num
		}
	}

	return result, false
}
