package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/markrtucker/adventofcode/2019/computer"
)

func main() {

	b, err := ioutil.ReadFile("/Users/marktucker/go/src/github.com/markrtucker/adventofcode/2019/day5/input.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(2)
	}

	s := string(b)
	ss := strings.Split(s, ",")
	orig := make([]int, len(ss))

	for i, sv := range ss {
		orig[i], _ = strconv.Atoi(sv)
	}

	var out bytes.Buffer
	ic := computer.Intcode{
		In:  os.Stdin,
		Out: &out,
	}
	ic.Compute(orig)

	// computer.Compute(orig)
	// computer.Compute([]int{3, 9, 5, 9, 6, 99, 4, 9, 99, -1, 8})
	fmt.Println("FINAL OUTPUT:", string(out.Bytes()))
}
