package main

import "testing"

func Test_gcd(t *testing.T) {

	tests := []struct {
		m   int
		n	int
		result int
	}{
		{15, 51, 3},
		{51, 15, 3},
	}

	for _, tt := range tests {
		got := gcd(tt.m, tt.n)
		if tt.result != got {
			t.Errorf("got %v, want %v", got, tt.result)
		}
	}
}
