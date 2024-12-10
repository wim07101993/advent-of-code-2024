package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	solution := Solve(strings.NewReader(input))

	if solution != 31 {
		t.Errorf("want %d, got %d", 4, solution)
	}
}
