package computer

import (
	"fmt"
	"os"
	"strconv"
)

var input = opcode{
	nParams: 1,
	op:      read,
	out:     noopOut,
	nextIP:  plus(2),
}

var output = opcode{
	nParams: 1,
	op:      echo,
	out:     noopOut,
	nextIP:  plus(2),
}

func echo(mem []int, target int, n ...int) int {
	// r := mem[n[0]]
	// r := mem[target] // TODO: is this right???
	bs := []byte(strconv.Itoa(n[0]))
	GlobalStdOut.Write(bs)
	// fmt.Println("OUTPUT: ", r)
	return n[0]
}

func read(mem []int, target int, n ...int) int {
	// return 9

	var i int
	_, err := fmt.Fscanf(GlobalStdIn, "%d\n", &i)
	if err != nil {
		fmt.Println("Error scanning input ", err)
		os.Exit(2)
	}
	mem[target] = i
	return i
}
