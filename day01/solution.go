package day01

import (
	"bufio"
	"io"
	"strconv"
)

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFuelWithOverhead(mass int) int {
	fuel := calculateFuel(mass)
	totalFuel := fuel

	for fuel > 0 {
		fuel = calculateFuel(fuel)

		if fuel > 0 {
			totalFuel += fuel
		}
	}

	return totalFuel
}

func SolvePart01(input []int) int {
	fuel := 0

	for _, mass := range input {
		fuel += calculateFuel(mass)
	}

	return fuel
}

func SolvePart02(input []int) int {
	fuel := 0

	for _, mass := range input {
		fuel += calculateFuelWithOverhead(mass)
	}

	return fuel
}

func Solve(part int, input io.Reader) int {
	masses := make([]int, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		masses = append(masses, mass)
	}

	if part == 1 {
		return SolvePart01(masses)
	}

	return SolvePart02(masses)
}
