package day13

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"aoc2019/intcode"
)

const (
	tileEmpty = iota
	tileWall
	tileBlock
	tilePaddle
	tileBall
)

type tile struct {
	x  int
	y  int
	id int
}

func SolvePart01(program []int) int {
	in := make(chan int, 1)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go func() {
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}()

	blockTileCount := 0

	for {
		_, more := <-out
		if !more {
			break
		}

		<-out
		id := <-out

		if id == tileBlock {
			blockTileCount++
		}
	}

	return blockTileCount
}

type grid struct {
	tiles [][]int
}

func newGrid() *grid {
	g := grid{}
	g.tiles = make([][]int, 1)
	g.tiles[0] = make([]int, 1)
	g.tiles[0][0] = tileEmpty

	return &g
}

func (g *grid) set(x, y, id int) {
	width, height := g.dim()

	for height <= y {
		g.tiles = append(g.tiles, make([]int, width))
		width, height = g.dim()
	}

	for width <= x {
		for i := 0; i < len(g.tiles); i++ {
			g.tiles[i] = append(g.tiles[i], 0)
		}

		width, height = g.dim()
	}

	g.tiles[y][x] = id
}

func (g *grid) get(x, y int) int {
	return g.tiles[y][x]
}

func (g *grid) dim() (int, int) {
	return len(g.tiles[0]), len(g.tiles)
}

func SolvePart02(program []int) int {
	program[0] = 2
	in := make(chan int)
	out := make(chan int)
	signal := make(chan int)

	c := intcode.NewComputer(program, in, out)
	c.AddSignalHandler(signal)
	g := newGrid()

	score := 0
	ballX := 0
	paddleX := 0

	go func() {
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}()

	for {
		s, more := <-signal
		if !more {
			break
		}

		switch s {
		case intcode.SignalIn:
			switch {
			case paddleX < ballX:
				in <- 1
			case paddleX > ballX:
				in <- -1
			default:
				in <- 0
			}
		case intcode.SignalOut:
			x := <-out
			<-signal
			y := <-out
			<-signal
			id := <-out

			if x == -1 && y == 0 {
				score = id
				continue
			}

			g.set(x, y, id)

			if id == tileBall {
				ballX = x
			}
			if id == tilePaddle {
				paddleX = x
			}
		}
	}

	return score
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
