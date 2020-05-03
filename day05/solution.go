package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"aoc2019/intcode"
)

func toIntcode(input []string) []int {
	intcode := make([]int, len(input))

	for i, v := range input {
		x, _ := strconv.Atoi(v)
		intcode[i] = x
	}

	return intcode
}

func SolvePart01(program []int) int {
	var builder strings.Builder
	c := intcode.NewComputer(program, strings.NewReader("1\n"), &builder)
	c.Run()

	lines := strings.Fields(builder.String())
	result, _ := strconv.Atoi(lines[len(lines)-1])
	return result
}

func SolvePart02(program []int) int {
	var builder strings.Builder
	c := intcode.NewComputer(program, strings.NewReader("5\n"), &builder)
	c.Run()

	lines := strings.Fields(builder.String())
	result, _ := strconv.Atoi(lines[len(lines)-1])
	return result
}

func Solve(part int, input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	values := strings.Split(scanner.Text(), ",")

	program := make([]int, len(values))
	for i, v := range values {
		num, _ := strconv.Atoi(v)
		program[i] = num
	}

	if part == 1 {
		return SolvePart01(program)
	}

	return SolvePart02(program)
}
