package day07

import (
	"aoc2019/intcode"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func toIntcode(input string) []int {
	segments := strings.Split(input, ",")
	intcode := make([]int, len(input))

	for i, v := range segments {
		x, _ := strconv.Atoi(v)
		intcode[i] = x
	}

	return intcode
}

func run(program []int, phases []int) int {
	inputs := make([]*io.PipeReader, len(phases)+1)
	outputs := make([]*io.PipeWriter, len(phases)+1)

	for i := 0; i < len(phases)+1; i++ {
		inputs[i], outputs[i] = io.Pipe()
	}

	for i, phase := range phases {
		input := inputs[i]
		output := outputs[i+1]
		c := intcode.NewComputer(program, input, output)

		go func() {
			err := c.Run()
			if err != nil {
				panic(fmt.Sprintf("while executing program: %v", err))
			}

			err = output.Close()
			if err != nil {
				panic(fmt.Sprintf("while closing pipe: %v", err))
			}
		}()

		fmt.Fprintln(outputs[i], phase)
	}

	input := outputs[0]
	output := inputs[len(phases)]

	var signal int
	fmt.Fprintln(input, 0)
	fmt.Fscanln(output, &signal)

	input.Close()
	return signal
}

func heap(k int, nums []int, process func([]int)) {
	if k == 1 {
		process(nums)
	} else {
		heap(k-1, nums, process)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}

			heap(k-1, nums, process)
		}
	}
}

func permutations(nums []int) [][]int {
	perms := make([][]int, 0)

	add := func(p []int) {
		c := make([]int, len(p))
		copy(c, p)
		perms = append(perms, c)
	}

	heap(len(nums), nums, add)

	return perms
}

func SolvePart01(program []int) int {
	options := permutations([]int{0, 1, 2, 3, 4})
	max := run(program, options[0])

	for i := 1; i < len(options); i++ {
		out := run(program, options[i])
		if out > max {
			max = out
		}
	}

	return max
}

func SolvePart02(program []int) int {
	return 0
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
