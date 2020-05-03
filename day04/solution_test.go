package day04

import (
	"reflect"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		number int
		digits []int
	}{
		{123456, []int{1, 2, 3, 4, 5, 6}},
		{111111, []int{1, 1, 1, 1, 1, 1}},
		{223450, []int{2, 2, 3, 4, 5, 0}},
		{102345, []int{1, 0, 2, 3, 4, 5}},
		{12345, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		ds := digits(test.number)
		if !reflect.DeepEqual(test.digits, ds) {
			t.Errorf("%d - expected %v, got %v", test.number, test.digits, ds)
		}
	}
}

func TestNeverDecreasing(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{123456, true},
		{111111, true},
		{223450, false},
		{102345, false},
	}

	for _, test := range tests {
		if neverDecreasing(test.number) != test.expected {
			t.Errorf("%d - expected %v", test.number, test.expected)
		}
	}
}

func TestHasDouble(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{123456, false},
		{112345, true},
		{111111, true},
		{123455, true},
	}

	for _, test := range tests {
		if hasDouble(test.number) != test.expected {
			t.Errorf("%d - expected %v", test.number, test.expected)
		}
	}
}

func TestHasBoundedDouble(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{123456, false},
		{112233, true},
		{123444, false},
		{111122, true},
		{123455, true},
	}

	for _, test := range tests {
		if hasBoundedDouble(test.number) != test.expected {
			t.Errorf("%d - expected %v", test.number, test.expected)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 1748},
		{2, 1180},
	}

	for _, test := range tests {
		got := Solve(test.part)
		if got != test.expected {
			t.Errorf("failed part %d: expected %d, got %d", test.part, test.expected, got)
		}
	}
}
