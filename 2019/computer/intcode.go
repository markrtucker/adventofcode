package computer

import (
	"fmt"
	"io"
	"math"
	"os"
)

// TODO Make intcode (computer) struct

// GlobalStdIn is a hack to allow unit tests to read input from "stdin"
var GlobalStdIn io.Reader = os.Stdin

// GlobalStdOut is a hack to allow unit tests to read input from "stdin"
var GlobalStdOut io.Writer = os.Stdout

// TODO - make this an interface?
type opcode struct {
	nParams int
	op      func(mem []int, target int, n ...int) int
	out     func(mem []int, result int, target int)
	nextIP  func(ip int) int
}

var opcodes = map[int]opcode{
	1:  add,
	2:  mult,
	3:  input,
	4:  output,
	5:  jumpIfTrue,
	6:  jumpIfFalse,
	7:  lessThan,
	8:  equal,
	99: halt,
}

// Intcode is a computer
type Intcode struct {
	In    io.Reader
	Out   io.Writer
	Debug bool
}

// Compute TODO godoc
func (i Intcode) Compute(mem []int) int {
	GlobalStdIn = i.In
	GlobalStdOut = i.Out
	ip := 0

	if i.Debug {
		fmt.Println("Init: ", mem)
	}

	for {
		inst, args, target := decode(mem, ip)

		if i.Debug {
			fmt.Println("Instruction: ", inst)
		}
		// fmt.Printf("Op: %v\n", inst.op)
		// out :=
		n := inst.op(mem, target, args...)
		if n == -99 {
			break
		}

		// inst.out(mem, out, target)

		ip = inst.nextIP(ip)
		if i.Debug {
			fmt.Println("Now: ", mem)
		}
	}

	return -99
}

func decodeOp(opCode int) (opcode, []int) {
	oc := opCode % 100
	op := opcodes[oc]
	addrModes := make([]int, op.nParams)

	for i := 0; i < op.nParams; i++ {
		div := int(math.Pow(10, float64(i+2)))
		addrModes[i] = opCode / div % 10
	}

	// fmt.Printf("Decoded %d: %v\n", oc, addrModes)
	return op, addrModes
}

func decode(mem []int, ip int) (opcode, []int, int) {

	// fmt.Printf("Decoding %d at ip %d\n", mem[ip], ip)
	op, addrModes := decodeOp(mem[ip])

	args := make([]int, op.nParams)
	for i := 0; i < op.nParams; i++ {
		args[i] = mem[ip+i+1]
		if addrModes[i] == 0 { //&& i < op.nParams-1 { // TODO: Last param is not resolved - not always true....???
			// resolve args - TODO: handle addressing modes
			args[i] = mem[args[i]]
		}
	}

	target := mem[ip+op.nParams] // TODO - this is wrong...just treat as arg...

	return op, args, target
}
