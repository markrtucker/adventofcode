package computer

var add = opcode{
	nParams: 3,
	op:      addOp,
	out:     noopOut,
	nextIP:  plus(4),
}

func addOp(mem []int, target int, n ...int) int {
	mem[target] = n[0] + n[1]
	return mem[target]
}

var mult = opcode{
	nParams: 3,
	op:      multOp,
	out:     noopOut,
	nextIP:  plus(4),
}

func multOp(mem []int, target int, n ...int) int {
	mem[target] = n[0] * n[1]
	return mem[target]
}

// func store(mem []int, r int, t int) {
// 	mem[t] = r
// }

func plus(inc int) func(int) int {
	return func(n int) int { return n + inc }
}
