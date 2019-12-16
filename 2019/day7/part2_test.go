package day7

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/markrtucker/adventofcode/2019/computer"
)

var part2Test1 = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`

func Test_Part2(t *testing.T) {

	ss := strings.Split(part2Test1, ",")
	prog := make([]int, len(ss))

	for i, sv := range ss {
		prog[i], _ = strconv.Atoi(sv)
	}

	s := []int{5, 6, 7, 8, 9}

	max := 0
	var maxSeq []int
	Perm(s, func(seq []int) {
		// fmt.Println(string(a))

		// for _, seq := range seqs {
		v := runPart2(prog, seq)
		fmt.Printf("Value is %d for sequence %v\n", v, seq)
		if v > max {
			max = v
			maxSeq = seq
		}
		// }

	})

	fmt.Println()
	fmt.Printf("MAX VALUE: %d for sequence %v\n", max, maxSeq)
	fmt.Println()
	t.Fail()

}

func runPart2(prog []int, seq []int) int {
	// in1 := os.Stdin
	var in1 bytes.Buffer
	var out1, out2, out3, out4 bytes.Buffer

	// TODO - this doesn't work. Thinking about potential switching to channels
	// to communicate between a1, a2, a3, ...

	// var c0, c1, c2, c3, c4 chan string

	fmt.Println("Seq", seq)

	in1.WriteString(strconv.Itoa(seq[0]))
	in1.WriteString("\n")
	in1.WriteString("0")
	out1.WriteString(strconv.Itoa(seq[1]))
	out1.WriteString("\n")
	out2.WriteString(strconv.Itoa(seq[2]))
	out2.WriteString("\n")
	out3.WriteString(strconv.Itoa(seq[3]))
	out3.WriteString("\n")
	out4.WriteString(strconv.Itoa(seq[4]))
	out4.WriteString("\n")

	a1 := computer.Intcode{
		In:    &in1,
		Out:   &out1,
		Debug: true,
	}

	a2 := computer.Intcode{
		In:  &out1,
		Out: &out2,
	}

	a3 := computer.Intcode{
		In:  &out2,
		Out: &out3,
	}

	a4 := computer.Intcode{
		In:  &out3,
		Out: &out4,
	}

	a5 := computer.Intcode{
		In:  &out4,
		Out: &in1,
	}

	// for {
	fmt.Println("Start")
	go a1.Compute(prog)
	fmt.Println("a1")
	go a2.Compute(prog)
	fmt.Println("a2")
	go a3.Compute(prog)
	fmt.Println("a3")
	go a4.Compute(prog)
	fmt.Println("a4")
	go a5.Compute(prog)
	val, _ := strconv.Atoi(string(in1.Bytes()))
	fmt.Println("Curr val is ", val)
	// }

	for i := 0; i <= 10; i++ {
		time.Sleep(1000)
	}
	// val, _ := strconv.Atoi(string(out5.Bytes()))
	return val
}

// var test1 = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
// var test2 = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
// var test3 = `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`

// var day7input = `3,8,1001,8,10,8,105,1,0,0,21,34,51,68,89,98,179,260,341,422,99999,3,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,1002,9,2,9,4,9,99,3,9,1001,9,3,9,102,3,9,9,101,4,9,9,4,9,99,3,9,102,2,9,9,101,2,9,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,99`
