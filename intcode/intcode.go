package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Computer contains the execution of an intcode program
type Computer struct {
	pc      int
	scanner *bufio.Scanner
	out     io.Writer
	Memory  []int
	Halted  bool
}

// Operations
const (
	OpAdd         = 1
	OpMultiply    = 2
	OpInput       = 3
	OpOutput      = 4
	OpJumpIfTrue  = 5
	OpJumpIfFalse = 6
	OpLessThan    = 7
	OpEquals      = 8
	OpHalt        = 99
)

// Parmeter modes
const (
	ModePosition  = 0
	ModeImmediate = 1
)

// Instruction represents the information contained in an opcode
type Instruction struct {
	Op    int
	Mode1 int
	Mode2 int
	Mode3 int
}

// Decode decodes an opcode into an Instruction
func Decode(opcode int) (Instruction, error) {
	op := opcode % 100

	if op < 1 {
		return Instruction{}, fmt.Errorf("illegal operation %d in opcode %d", op, opcode)
	}

	if op > 8 && op != 99 {
		return Instruction{}, fmt.Errorf("illegal operation %d in opcode %d", op, opcode)
	}

	mode1 := (opcode / 100) % 10
	if mode1 != ModePosition && mode1 != ModeImmediate {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 1 in opcode %d", mode1, opcode)
	}

	mode2 := (opcode / 1000) % 10
	if mode2 != ModePosition && mode2 != ModeImmediate {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 2 in opcode %d", mode2, opcode)
	}

	mode3 := (opcode / 10000) % 10
	if mode3 != ModePosition && mode3 != ModeImmediate {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 3 in opcode %d", mode3, opcode)
	}

	return Instruction{op, mode1, mode2, mode3}, nil
}

func (c *Computer) readInd(addr int) int {
	return c.Memory[c.Memory[addr]]
}

func (c *Computer) writeInd(addr int, value int) {
	c.Memory[c.Memory[addr]] = value
}

func (c *Computer) fetchParam(offset, mode int) int {
	if mode == ModePosition {
		return c.readInd(c.pc + offset)
	}

	return c.Memory[c.pc+offset]
}

func (c *Computer) step() error {
	op, err := Decode(c.Memory[c.pc])
	if err != nil {
		return fmt.Errorf("failed decode at address %d: %v", c.pc, err)
	}

	switch op.Op {
	case OpAdd:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)
		c.writeInd(c.pc+3, x+y)
		c.pc += 4
	case OpMultiply:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)
		c.writeInd(c.pc+3, x*y)
		c.pc += 4
	case OpInput:
		if !c.scanner.Scan() {
			err := c.scanner.Err()
			if err != nil {
				return fmt.Errorf("unexpected error while reading input: %v", err)
			}

			return fmt.Errorf("unexpected end of input")
		}

		raw := c.scanner.Text()
		v, err := strconv.Atoi(raw)
		if err != nil {
			return fmt.Errorf("unexpected input: %v", err)
		}

		c.writeInd(c.pc+1, v)
		c.pc += 2
	case OpOutput:
		v := c.fetchParam(1, op.Mode1)
		_, err := fmt.Fprintln(c.out, v)
		if err != nil {
			return err
		}
		c.pc += 2
	case OpJumpIfTrue:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)

		if x == 0 {
			c.pc += 3
		} else {
			c.pc = y
		}
	case OpJumpIfFalse:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)

		if x == 0 {
			c.pc = y
		} else {
			c.pc += 3
		}
	case OpLessThan:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)

		if x < y {
			c.writeInd(c.pc+3, 1)
		} else {
			c.writeInd(c.pc+3, 0)
		}

		c.pc += 4
	case OpEquals:
		x := c.fetchParam(1, op.Mode1)
		y := c.fetchParam(2, op.Mode2)

		if x == y {
			c.writeInd(c.pc+3, 1)
		} else {
			c.writeInd(c.pc+3, 0)
		}

		c.pc += 4
	case OpHalt:
		c.Halted = true
	}

	return nil
}

// NewComputer returns a new intcode computer intialized from program
func NewComputer(program []int, r io.Reader, w io.Writer) Computer {
	c := Computer{
		Memory:  make([]int, len(program)),
		scanner: bufio.NewScanner(r),
		out:     w,
	}

	copy(c.Memory, program)

	return c
}

// Run runs the program until halted
func (c *Computer) Run() error {
	for !c.Halted {
		if err := c.step(); err != nil {
			return err
		}
	}

	return nil
}