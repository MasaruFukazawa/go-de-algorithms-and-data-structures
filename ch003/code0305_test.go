package main

import "testing"

func TestSubsetSumBitBruteForce(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"Normal case", []int{1, 2, 4, 5, 11}, 10, true},
		{"Not found", []int{1, 2, 4, 5, 11}, 100, false},
		{"Empty array k=0", []int{}, 0, true},
		{"Single match", []int{5}, 5, true},
		{"Single no match", []int{5}, 3, false},
		{"Sum of all", []int{1, 2, 3}, 6, true},
		{"k=0 with elements", []int{1, 2, 3}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetSumBitBruteForce(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("subsetSumBitBruteForce(%v, %d) = %v, want %v",
					tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
