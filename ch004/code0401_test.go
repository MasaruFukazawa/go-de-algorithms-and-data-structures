package main

import "testing"

func Test_recursiveAddFunc(t *testing.T) {

	tests := []struct {
		nums   int
		result int
	}{
		{10, 55},
	}

	for _, tt := range tests {
		got := recursiveAddFunc(tt.nums)
		if tt.result != got {
			t.Errorf("got %v, want %v", got, tt.result)
		}
	}
}
