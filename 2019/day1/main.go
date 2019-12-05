package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("Hello")

	b, err := ioutil.ReadFile("/Users/marktucker/go/src/github.com/markrtucker/adventofcode/2019/day1/input.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	total := 0
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Error opening file: ", err)
		}

		fuel := calcFuelRecursive(v)
		total += fuel
	}

	fmt.Println(total)
}

func calcFuelRecursive(m int) int {

	total := 0

	f := calcFuel(m)
	for f > 0 {
		total += f
		f = calcFuel(f)
	}

	return total
}

func calcFuel(m int) int {
	n := m / 3
	return n - 2
}
