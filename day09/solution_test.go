package day09

import (
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 4080871669},
		{2, 75202},
	}

	for _, test := range tests {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		got := Solve(test.part, file)
		if got != test.expected {
			t.Errorf("failed part %d: expected %d, got %d", test.part, test.expected, got)
		}

		file.Close()
	}
}
