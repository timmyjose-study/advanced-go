package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// tests have to begin with "Test"
// also, the files themselves have to be named <file-name>_test.go
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("Expected %d, found %d\n", -2, ans)
	}
}

// table-driven testing
func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{0, 1, 0},
		{1, 2, 1},
		{2, -2, -2},
		{100, 9, 9},
		{-11, -89, -89}, // why is this comma needed?
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) { // t.Run runs each test individually
			ans := IntMin(tt.a, tt.b) //
			if ans != tt.expected {
				t.Errorf("expected %d, got %d\n", tt.expected, ans)
			}
		})
	}
}
