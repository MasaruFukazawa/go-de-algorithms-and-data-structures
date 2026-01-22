package main

import "testing"

func Test_fibo(t *testing.T) {

	tests := []struct {
		n	int
		result int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, tt := range tests {
		got := fibo(tt.n)
		if tt.result != got {
			t.Errorf("got %v, want %v", got, tt.result)
		}
	}
}
