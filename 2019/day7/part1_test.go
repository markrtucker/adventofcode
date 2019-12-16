package day7

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/markrtucker/adventofcode/2019/computer"
)

func Test_Something(t *testing.T) {

	if true {
		return
	}

	ss := strings.Split(day7input, ",")
	prog := make([]int, len(ss))

	for i, sv := range ss {
		prog[i], _ = strconv.Atoi(sv)
	}

	s := []int{0, 1, 2, 3, 4}

	max := 0
	var maxSeq []int
	Perm(s, func(seq []int) {
		// fmt.Println(string(a))

		// for _, seq := range seqs {
		v := runIt(prog, seq)
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

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func runIt(prog []int, seq []int) int {
	// in1 := os.Stdin
	var in1 bytes.Buffer
	var out1, out2, out3, out4, out5 bytes.Buffer

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
		In:  &in1,
		Out: &out1,
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
		Out: &out5,
	}

	a1.Compute(prog)
	a2.Compute(prog)
	a3.Compute(prog)
	a4.Compute(prog)
	a5.Compute(prog)

	val, _ := strconv.Atoi(string(out5.Bytes()))
	return val
}

var test1 = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
var test2 = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
var test3 = `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`

var day7input = `3,8,1001,8,10,8,105,1,0,0,21,34,51,68,89,98,179,260,341,422,99999,3,9,1001,9,4,9,102,4,9,9,4,9,99,3,9,1002,9,5,9,1001,9,2,9,1002,9,2,9,4,9,99,3,9,1001,9,3,9,102,3,9,9,101,4,9,9,4,9,99,3,9,102,2,9,9,101,2,9,9,1002,9,5,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,101,2,9,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,99`
