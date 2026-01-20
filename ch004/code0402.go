package main

import "fmt"

func recursiveAddFunc2(num int) int {

	fmt.Printf("%dを呼び出しました\n", num)

	if num <= 0 {
		return 0
	}

	result := num + recursiveAddFunc2(num-1)
	fmt.Printf("%dまでの和 : %d\n", num, result)

	return result
}
