package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	b, err := ioutil.ReadFile("/Users/marktucker/go/src/github.com/markrtucker/adventofcode/2019/day3/input.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(2)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	s.Scan()
	l1 := s.Text()
	s.Scan()
	l2 := s.Text()

	dist := FindIntersection(l1, l2)

	fmt.Println(dist)
}

// FindIntersection todo godoc
func FindIntersection(l1 string, l2 string) int {

	var m map[int]map[int]bool
	m = make(map[int]map[int]bool, 0)

	populateMap(m, l1)
	return -99
}

func populateMap(m map[int]map[int]bool, l1 string) {
	var x, y int
	commands := strings.Split(l1, ",")
	for _, c := range commands {
		n, _ := strconv.Atoi(c[1:])
		switch c[0] {
		case 'U':
			for i := 0; i < n; i++ {
				set(m, x, y)
				y++
			}
		case 'L':
			for i := 0; i < n; i++ {
				set(m, x, y)
				x--
			}
		case 'R':
			for i := 0; i < n; i++ {
				set(m, x, y)
				x++
			}
		case 'D':
			for i := 0; i < n; i++ {
				set(m, x, y)
				y--
			}
		default:
			fmt.Println("Unknown command " + c)
			os.Exit(3)
		}
	}
}

func set(m map[int]map[int]bool, x int, y int) {
	submap := m[x]
	if submap == nil {
		submap = make(map[int]bool, 0)
		m[x] = submap
	}
	m[x][y] = true
}
