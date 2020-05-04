package day08

import (
	"os"
	"reflect"
	"testing"
)

func TestSolvePart01(t *testing.T) {
	tests := []struct {
		digits   []int
		width    int
		height   int
		expected int
	}{
		{[]int{1, 1, 2, 2, 5, 6, 7, 8, 9, 0, 1, 2}, 3, 2, 4},
	}

	for _, test := range tests {
		result := SolvePart01(test.digits, test.width, test.height)

		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}

func TestSolvePart02(t *testing.T) {
	tests := []struct {
		digits   []int
		width    int
		height   int
		expected []int
	}{
		{[]int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}, 2, 2, []int{0, 1, 1, 0}},
	}

	for _, test := range tests {
		result := SolvePart02(test.digits, test.width, test.height)

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 1452},
		{2, 0},
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
