package day09

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"aoc2019/intcode"
)

func SolvePart01(program []int) int {
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go c.Run()

	in <- 1

	var result int

	for {
		r, more := <-out
		if !more {
			break
		}
		result = r
	}

	return result
}

func SolvePart02(program []int) int {
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go c.Run()
	in <- 2
	return <-out
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
