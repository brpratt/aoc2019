package day03

import (
	"os"
	"testing"
)

func TestParseSegment(t *testing.T) {
	table := []struct {
		input    string
		expected segment
	}{
		{"R8", segment{right, 8}},
		{"U5", segment{up, 5}},
		{"L5", segment{left, 5}},
		{"D3", segment{down, 3}},
	}

	for _, test := range table {
		segment := parseSegment(test.input)

		if segment != test.expected {
			t.Errorf("expected: %v, actual: %v", test.expected, segment)
		}
	}
}

func TestParsePath(t *testing.T) {
	input := "R8,U5,L5,D3"

	expected := []segment{
		segment{right, 8},
		segment{up, 5},
		segment{left, 5},
		segment{down, 3},
	}

	path := parsePath(input)

	if len(path) != len(expected) {
		t.Fatalf("expected: %v, actual: %v", expected, path)
	}

	for i := range path {
		if path[i] != expected[i] {
			t.Fatalf("expected: %v, actual: %v", expected, path)
		}
	}
}

func TestSegmentPoints(t *testing.T) {
	table := []struct {
		o              point
		s              segment
		expectedPoint  point
		expectedPoints []point
	}{
		{point{0, 0}, segment{up, 2}, point{0, 2}, []point{point{0, 1}, point{0, 2}}},
		{point{1, 1}, segment{left, 1}, point{0, 1}, []point{point{0, 1}}},
		{point{-3, -3}, segment{right, 2}, point{-1, -3}, []point{point{-2, -3}, point{-1, -3}}},
		{point{3, 4}, segment{down, 1}, point{3, 3}, []point{point{3, 3}}},
	}

	for _, test := range table {
		point, points := segmentPoints(test.o, test.s)

		if point != test.expectedPoint {
			t.Errorf("mismatched point: expected: %v, actual: %v", test.expectedPoint, point)
		}

		if len(points) == len(test.expectedPoints) {
			for i := range points {
				if points[i] != test.expectedPoints[i] {
					t.Errorf("mismatched points: expected: %v, actual: %v", test.expectedPoints, points)
				}
			}
		} else {
			t.Errorf("mismatched points: expected: %v, actual: %v", test.expectedPoints, points)
		}
	}
}

func TestPoints(t *testing.T) {
	o := point{0, 0}
	p := path{segment{up, 1}, segment{right, 2}}
	expectedPoints := []point{point{0, 1}, point{1, 1}, point{2, 1}}

	points := points(o, p)

	if len(points) != len(expectedPoints) {
		t.Fatalf("expected: %v, actual: %v", expectedPoints, points)
	}

	for key := range points {
		if points[key] != expectedPoints[key] {
			t.Fatalf("expected: %v, actual: %v", expectedPoints, points)
		}
	}
}

func TestInterset(t *testing.T) {
	pts1 := map[point]int{point{0, 0}: 1, point{1, 2}: 2, point{3, 3}: 3}
	pts2 := map[point]int{point{2, 1}: 1, point{0, 0}: 2, point{3, 3}: 3}
	expectedPts := map[point]bool{point{0, 0}: true, point{3, 3}: true}

	pts := intersect(pts1, pts2)

	if len(pts) != len(expectedPts) {
		t.Fatalf("expected: %v, actual: %v", expectedPts, pts)
	}

	for key := range pts {
		if pts[key] != expectedPts[key] {
			t.Fatalf("expected: %v, actual: %v", expectedPts, pts)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 280},
		{2, 10554},
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
