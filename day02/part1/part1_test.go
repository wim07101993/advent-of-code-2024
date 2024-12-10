package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
7 7 5 3 1
15 7 5 3 1
1 3 6 7 9`

	solution := Solve(strings.NewReader(input))

	if solution != 2 {
		t.Errorf("want %d, got %d", 4, solution)
	}
}
