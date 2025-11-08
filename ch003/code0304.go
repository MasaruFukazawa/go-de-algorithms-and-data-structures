package main

import "math"

func searchMinValidPairSum(nums1 []int, nums2 []int, k int) (int, bool) {

	sum_value := math.MaxInt
	result := math.MaxInt

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			sum_value = num1 + num2
			if sum_value >= k && sum_value < result {
				result = sum_value
			}
		}
	}

	if result ==  math.MaxInt {
		return math.MaxInt, true
	}

	return result, false
}
