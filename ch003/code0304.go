package main

import "math"

func searchMinValidPairSum(nums1 []int, nums2 []int, k int) (int, bool) {

	min_val := math.MaxInt
	found := false

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			sum_value := num1 + num2
			if sum_value >= k {
				if !found {
					min_val = sum_value
					found = true
				} else {
					if sum_value < min_val {
						min_val = sum_value
					}
				}
			}
		}
	}

	if !found {
		return math.MaxInt, true // No valid pair found
	}

	return min_val, false
}
