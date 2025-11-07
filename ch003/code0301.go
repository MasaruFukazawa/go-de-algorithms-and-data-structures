package main

func linerSearch(nums []int, target int) bool {

	result := false

	for _, num := range nums {
		if num == target {
			result = true
			break
		}
	}

	return result
}
