package computer

import (
	"reflect"
	"runtime"
)

//
// lessThan:
// if the first parameter is less than the second parameter, it stores 1
// in the position given by the third parameter. Otherwise, it stores 0.
//
var lessThan = opcode{
	nParams: 3,
	op:      ltOp,
	out:     noopOut,
	nextIP:  plus(4),
}

func ltOp(mem []int, target int, args ...int) int {

	if args[0] < args[1] {
		mem[target] = 1
		return 1
	}

	mem[target] = 0
	return 0
}

//
// equal
//
var equal = opcode{
	nParams: 3,
	op:      eqOp,
	out:     noopOut,
	nextIP:  plus(4),
}

func (o opcode) String() string {
	opName := runtime.FuncForPC(reflect.ValueOf(o.op).Pointer()).Name()

	return opName
}

func eqOp(mem []int, target int, args ...int) int {

	if args[0] == args[1] {
		mem[target] = 1
		return 1
	}

	mem[target] = 0
	return 0
}
