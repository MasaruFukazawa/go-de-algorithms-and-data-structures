package main

import "fmt"

func gcd(m int, n int) int {

	fmt.Printf("(%d, %d)を呼び出しました\n", m, n)

	if n == 0 {
		return m
	}

	return gcd(n, m % n)
}
