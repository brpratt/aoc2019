package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"aoc2019/day01"
	"aoc2019/day02"
	"aoc2019/day04"
	"aoc2019/day05"
	"aoc2019/day06"
	"aoc2019/day07"
	"aoc2019/day08"
	"aoc2019/day09"
	"aoc2019/day10"
	"aoc2019/day11"
	"aoc2019/day12"
	"aoc2019/day13"
	"aoc2019/day14"
	"aoc2019/day15"
)

type solution = func(int)

var solutions = []solution{
	solveDay01,
	solveDay02,
	solveDay03,
	solveDay04,
	solveDay05,
	solveDay06,
	solveDay07,
	solveDay08,
	solveDay09,
	solveDay10,
	solveDay11,
	solveDay12,
	solveDay13,
	solveDay14,
	solveDay15,
}

func withFile(solver func(int, io.Reader) int, part int, filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("unable to open file %s: %v", filename, err))
	}

	defer file.Close()

	return solver(part, file)
}

func solveDay01(part int) {
	fmt.Println(withFile(day01.Solve, part, "day01/input.txt"))
}

func solveDay02(part int) {
	fmt.Println(withFile(day02.Solve, part, "day02/input.txt"))
}

func solveDay03(part int) {
	fmt.Println(withFile(day02.Solve, part, "day03/input.txt"))
}

func solveDay04(part int) {
	fmt.Println(day04.Solve(part))
}

func solveDay05(part int) {
	fmt.Println(withFile(day05.Solve, part, "day05/input.txt"))
}

func solveDay06(part int) {
	fmt.Println(withFile(day06.Solve, part, "day06/input.txt"))
}

func solveDay07(part int) {
	fmt.Println(withFile(day07.Solve, part, "day07/input.txt"))
}

func solveDay08(part int) {
	fmt.Println(withFile(day08.Solve, part, "day08/input.txt"))
}

func solveDay09(part int) {
	fmt.Println(withFile(day09.Solve, part, "day09/input.txt"))
}

func solveDay10(part int) {
	fmt.Println(withFile(day10.Solve, part, "day10/input.txt"))
}

func solveDay11(part int) {
	fmt.Println(withFile(day11.Solve, part, "day11/input.txt"))
}

func solveDay12(part int) {
	fmt.Println(day12.Solve(part))
}

func solveDay13(part int) {
	fmt.Println(withFile(day13.Solve, part, "day13/input.txt"))
}

func solveDay14(part int) {
	fmt.Println(withFile(day14.Solve, part, "day14/input.txt"))
}

func solveDay15(part int) {
	fmt.Println(withFile(day15.Solve, part, "day15/input.txt"))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: aoc2019 <day> <part>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("day must be a number")
		os.Exit(1)
	}

	if day > len(solutions) {
		fmt.Printf("day must be less than %d\n", len(solutions)+1)
		os.Exit(1)
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("part must be a number")
		os.Exit(1)
	}

	if part != 1 && part != 2 {
		fmt.Println("part must be the value 1 or 2")
		os.Exit(1)
	}

	solutions[day-1](part)
}
