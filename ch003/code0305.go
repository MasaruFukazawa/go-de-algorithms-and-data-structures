package main


func subsetSumBitBruteForce(nums []int, k int) bool {
	length := len(nums)

	for bit := 0; bit < (1 << length); bit++ {
		sum := 0
		for i := 0; i < length; i++ {
			if bit&(1<<i) != 0 {
				sum += nums[i]
			}
		}
		if sum == k {
			return true
		}
	}

	return false
}
