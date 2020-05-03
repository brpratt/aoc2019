package day02

import (
	"bufio"
	"io"
	"os"
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
	c := intcode.NewComputer(program, os.Stdin, os.Stderr)
	c.Memory[1] = 12
	c.Memory[2] = 2
	c.Run()

	return c.Memory[0]
}

func SolvePart02(program []int) int {
	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			c := intcode.NewComputer(program, os.Stdin, os.Stderr)
			c.Memory[1] = noun
			c.Memory[2] = verb
			c.Run()

			if c.Memory[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}

	panic("could not find solution")
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
