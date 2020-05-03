package day06

import (
	"bufio"
	"io"
	"strings"
)

func buildOrbitMap(input []string) map[string]string {
	m := make(map[string]string)

	for _, line := range input {
		segments := strings.Split(line, ")")
		m[segments[1]] = segments[0]
	}

	return m
}

func buildReverseOrbitMap(to string, m map[string]string) map[string]string {
	revM := make(map[string]string)
	next := m[to]

	for next != "COM" {
		revM[next] = to
		to = next
		next = m[next]
	}

	revM[next] = to

	return revM
}

func countOrbits(obj string, m map[string]string) int {
	count := 0

	for obj != "COM" {
		obj = m[obj]
		count++
	}

	return count
}

func countTransfers(from string, to string, m map[string]string) int {
	rev := buildReverseOrbitMap(to, m)
	curr := m[from]
	count := 0

	for _, ok := rev[curr]; !ok; _, ok = rev[curr] {
		curr = m[curr]
		count++
	}

	for rev[curr] != to {
		curr = rev[curr]
		count++
	}

	return count
}

func SolvePart01(orbits []string) int {
	m := buildOrbitMap(orbits)
	count := 0

	for key := range m {
		count += countOrbits(key, m)
	}

	return count
}

func SolvePart02(orbits []string) int {
	m := buildOrbitMap(orbits)

	return countTransfers("YOU", "SAN", m)
}

func Solve(part int, input io.Reader) int {
	orbits := make([]string, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		orbits = append(orbits, scanner.Text())
	}

	if part == 1 {
		return SolvePart01(orbits)
	}

	return SolvePart02(orbits)
}
