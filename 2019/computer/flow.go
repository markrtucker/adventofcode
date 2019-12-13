package computer

import (
	"fmt"
)

var halt = opcode{
	nParams: 0,
	op:      printZeroAndExit,
	out:     noopOut,
	nextIP:  nil, // Program terminates
}

func printZeroAndExit(mem []int, target int, args ...int) int {
	fmt.Println("Halting with final value ", mem[0])
	// os.Exit(1)
	return -99
}

// func printZeroAndExit(mem []int, r int, t int) {
// 	fmt.Println("Halting with final value ", mem[0])
// 	os.Exit(1)
// }

//
// Jump-if-false
//
var jumpIfFalse = opcode{
	nParams: 2,
	op:      jifOp,
	out:     noopOut,
	nextIP:  jmpNextIP,
}

var next int // ACK - fix global variable

func jifOp(mem []int, target int, args ...int) int {
	ret := 0
	next = -1

	if args[0] == 0 {
		// jump
		ret = args[1]
		next = args[1]
	}

	return ret
}

func noopOut(mem []int, r int, t int) {
	// noop
}

func jmpNextIP(ip int) int {
	if next != -1 {
		return next
	}

	return ip + 3
}

//
// Jump-if-true
// if the first parameter is non-zero, it sets the instruction pointer to the
// value from the second parameter. Otherwise, it does nothing.
//
var jumpIfTrue = opcode{
	nParams: 2,
	op:      jitOp,
	out:     noopOut,
	nextIP:  jmpNextIP,
}

func jitOp(mem []int, target int, args ...int) int {
	ret := 0
	next = -1

	if args[0] != 0 {
		// jump
		ret = args[1]
		next = args[1]
	}

	return ret
}
