package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	b, err := ioutil.ReadFile("/Users/marktucker/go/src/github.com/markrtucker/adventofcode/2019/day2/input.txt")

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

	fmt.Println(orig)

	work := make([]int, len(orig))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(work, orig)
			work[1] = noun
			work[2] = verb

			val := compute(work)
			if val == 19690720 {
				ans := 100*noun + verb
				fmt.Printf("Found answer %d with noun %d and verb %d!\n", ans, noun, verb)
				os.Exit(0)
			}
		}
	}
}

func compute(mem []int) int {
	ip := 0

	for {
		switch mem[ip] {
		case 1:
			result := mem[mem[ip+1]] + mem[mem[ip+2]]
			mem[mem[ip+3]] = result
		case 2:
			result := mem[mem[ip+1]] * mem[mem[ip+2]]
			mem[mem[ip+3]] = result
		case 99:
			return mem[0]
		default:
			fmt.Printf("Unkown instruction code %d at IP %d\n", mem[ip], ip)
			os.Exit(1)
		}

		ip += 4
	}
}
