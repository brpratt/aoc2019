package day08

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func SolvePart01(digits []int, width int, height int) int {
	minZeroCount := width*height + 1
	oneCount := 0
	twoCount := 0
	index := 0

	for i := 0; i < len(digits)/(width*height); i++ {
		var zc int
		var oc int
		var tc int

		for j := 0; j < width*height; j++ {
			switch digits[index] {
			case 0:
				zc++
			case 1:
				oc++
			case 2:
				tc++
			}

			index++
		}

		if zc < minZeroCount {
			minZeroCount = zc
			oneCount = oc
			twoCount = tc
		}
	}

	return oneCount * twoCount
}

func merge(layers [][]int) []int {
	final := make([]int, len(layers[0]))

	for i := 0; i < len(final); i++ {
		for j := 0; j < len(layers); j++ {
			if layers[j][i] != 2 {
				final[i] = layers[j][i]
				break
			}
		}
	}

	return final
}

func SolvePart02(digits []int, width int, height int) []int {
	var index int
	layers := make([][]int, len(digits)/(width*height))

	for i := range layers {
		layers[i] = make([]int, width*height)
		for j := 0; j < width*height; j++ {
			layers[i][j] = digits[index]
			index++
		}
	}

	return merge(layers)
}

func Solve(part int, input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()

	digits := make([]int, len(line))

	for i, v := range strings.Split(line, "") {
		d, _ := strconv.Atoi(v)
		digits[i] = d
	}

	if part == 1 {
		return SolvePart01(digits, 25, 6)
	}

	image := SolvePart02(digits, 25, 6)
	var index int

	for i := 0; i < 6; i++ {
		for j := 0; j < 25; j++ {
			if image[index] == 1 {
				fmt.Print("+")
			} else {
				fmt.Print(" ")
			}
			index++
		}
		fmt.Println()
	}

	return 0
}
