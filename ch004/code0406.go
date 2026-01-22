package main

import "fmt"

func fibo2(n int) int {

	fmt.Printf("%dを呼び出しました\n", n)

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := fibo2(n - 1) + fibo(n - 2)

	fmt.Printf("%d 項目 : %d\n", n, result)

	return result
}
