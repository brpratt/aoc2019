package day01

import (
	"os"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		fuel := calculateFuel(table.mass)
		if fuel != table.fuel {
			t.Errorf("mass %d: [expected %d] [actual %d]", table.mass, table.fuel, fuel)
		}
	}
}

func TestSolvePart01(t *testing.T) {
	input := []int{12, 14, 1969, 100756}
	expected := 34241

	value := SolvePart01(input)
	if value != expected {
		t.Fatalf("[expected %d] [actual %d]", expected, value)
	}
}

func TestCalculateFuelWithOverhead(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		fuel := calculateFuelWithOverhead(table.mass)
		if fuel != table.fuel {
			t.Errorf("mass %d: [expected %d] [actual %d]", table.mass, table.fuel, fuel)
		}
	}
}

func TestSolvePart02(t *testing.T) {
	input := []int{14, 1969, 100756}
	expected := 51314

	value := SolvePart02(input)
	if value != expected {
		t.Fatalf("[expected %d] [actual %d]", expected, value)
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 3317668},
		{2, 4973628},
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
