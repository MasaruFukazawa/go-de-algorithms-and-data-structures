package main

func recursiveAddFunc(num int) int {

	if num <= 0 {
		return 0
	}

	return num + recursiveAddFunc(num-1)
}
