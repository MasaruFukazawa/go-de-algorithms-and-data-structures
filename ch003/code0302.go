package main

func linerSearch02(nums []int, target int) []int {

	result := []int{}

	for key, num := range nums {
		if num == target {
			result = append(result, key)
		}
	}

	return result
}
