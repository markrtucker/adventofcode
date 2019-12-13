package computer

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var computeTests = []struct {
	name       string
	input      string
	program    []int
	expOut     []byte
	finalState []int
}{
	{"Simple-add", "", []int{1, 0, 0, 0, 99}, nil, []int{2, 0, 0, 0, 99}},
	{"Simple-mult", "", []int{2, 3, 0, 3, 99}, nil, []int{2, 3, 0, 6, 99}},
	{"Another-mult", "", []int{2, 4, 4, 5, 99, 0}, nil, []int{2, 4, 4, 5, 99, 9801}},
	{"More-math", "", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, nil, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	{"Output-99", "", []int{4, 2, 99}, []byte("99"), nil},
	{"Output-999", "", []int{104, 999, 99}, []byte("999"), nil},
	{"Simple-echo", "4", []int{3, 0, 4, 0, 99}, []byte("4"), nil},
	{"io-pos-equal-false", "7", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []byte("0"), nil},
	{"io-pos-equal-true", "8", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []byte("1"), nil},
	{"io-pos-lt-false", "88", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []byte("0"), nil},
	{"io-pos-lt-true", "5", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []byte("1"), nil},
	{"io-imm-equal-false", "7", []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []byte("0"), nil},
	{"io-imm-equal-true", "8", []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []byte("1"), nil},
	{"io-imm-lt-false", "9", []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []byte("0"), nil},
	{"io-imm-lt-true", "7", []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []byte("1"), nil},
	{"out-999-lt-8", "7", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []byte("999"), nil},
	{"out-1000-eq-8", "8", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []byte("1000"), nil},
	{"out-1001-gt-8", "9", []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []byte("1001"), nil},
}

func TestComputer(t *testing.T) {
	for idx, tt := range computeTests {
		fmt.Println("Test ", idx)
		doComputerTest(t, tt.name, tt.input, tt.program, tt.expOut, tt.finalState)
	}
}

func TestDebug(t *testing.T) {
	// Allows easy debugging of one particular test case, without commenting stuff out...
	// Set the appropriate index....
	indexToTest := 4

	// This runs the test...
	tt := computeTests[indexToTest]
	doComputerTest(t, tt.name, tt.input, tt.program, tt.expOut, tt.finalState)
}

func doComputerTest(t *testing.T, name string, input string, prog []int, expOut []byte, expFinal []int) {
	var buf bytes.Buffer
	ic := Intcode{
		In:  strings.NewReader(input),
		Out: &buf,
	}
	ic.Compute(prog)

	fmt.Println("Final output: ", buf.Bytes())

	if expOut != nil {
		assert.Equal(t, expOut, buf.Bytes(), "Wrong output: %s", name)
	}

	if expFinal != nil {
		assert.Equal(t, expFinal, prog, "Wrong final state: %s", name)
	}
}

// func TestSomething(t *testing.T) {
// 	input := "7" //\n"

// 	// testIfLessThenOrGreaterThan8
// 	prog := []int{4, 2, 99}

// 	// take one input, compare it to the value 8, and then produce one output:
// 	// prog := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8} // position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
// 	// prog := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8} // position mode, consider whether the input is less than to 8; output 1 (if it is) or 0 (if it is not).

// 	var buf bytes.Buffer
// 	ic := Intcode{
// 		In:  strings.NewReader(input),
// 		Out: &buf,
// 	}
// 	ic.Compute(prog)

// 	fmt.Println("Final output: ", buf.Bytes())
// 	assert.Equal(t, byte(99), buf.Bytes()[0])
// }

// func TestSomething(t *testing.T) {
// 	input := "7\n"

// 	// testIfLessThenOrGreaterThan8
// 	prog := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

// 	// take one input, compare it to the value 8, and then produce one output:
// 	// prog := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8} // position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
// 	// prog := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8} // position mode, consider whether the input is less than to 8; output 1 (if it is) or 0 (if it is not).

// 	var buf bytes.Buffer
// 	ic := Intcode{
// 		In:  strings.NewReader(input),
// 		Out: &buf,
// 	}
// 	ic.Compute(prog)

// 	fmt.Println(buf.Bytes())

// 	// We never get here - the 99 opcode calls os.Exit(...)
// 	t.Fail()
// }
