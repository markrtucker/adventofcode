package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	// var m map[int]map[int]int
	m := make(map[int]map[int]int, 0)

	populateMap(m, l1)

	dists := findIntersections(m, l2)

	fmt.Println(dists)

	sort.Ints(dists)

	return dists[0]
}

func findIntersections(m map[int]map[int]int, l1 string) []int {
	dists := make([]int, 0)

	var x, y int
	count := 0
	commands := strings.Split(l1, ",")
	for _, c := range commands {
		n, _ := strconv.Atoi(c[1:])
		switch c[0] {
		case 'U':
			for i := 0; i < n; i++ {
				if isSet(m, x, y) {
					dists = append(dists, computeDist(m, x, y, count))
				}
				count++
				y++
			}
		case 'L':
			for i := 0; i < n; i++ {
				if isSet(m, x, y) {
					dists = append(dists, computeDist(m, x, y, count))
				}
				count++
				x--
			}
		case 'R':
			for i := 0; i < n; i++ {
				if isSet(m, x, y) {
					dists = append(dists, computeDist(m, x, y, count))
				}
				count++
				x++
			}
		case 'D':
			for i := 0; i < n; i++ {
				if isSet(m, x, y) {
					dists = append(dists, computeDist(m, x, y, count))
				}
				count++
				y--
			}
		default:
			fmt.Println("Unknown command " + c)
			os.Exit(3)
		}
	}

	return dists
}

func computeDist(m map[int]map[int]int, x, y int, count int) int {

	return m[x][y] + count
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func populateMap(m map[int]map[int]int, l1 string) {
	var x, y int
	count := 0
	commands := strings.Split(l1, ",")
	for _, c := range commands {
		n, _ := strconv.Atoi(c[1:])
		switch c[0] {
		case 'U':
			for i := 0; i < n; i++ {
				set(m, x, y, count)
				count++
				y++
			}
		case 'L':
			for i := 0; i < n; i++ {
				set(m, x, y, count)
				count++
				x--
			}
		case 'R':
			for i := 0; i < n; i++ {
				set(m, x, y, count)
				count++
				x++
			}
		case 'D':
			for i := 0; i < n; i++ {
				set(m, x, y, count)
				count++
				y--
			}
		default:
			fmt.Println("Unknown command " + c)
			os.Exit(3)
		}
	}
}

func set(m map[int]map[int]int, x int, y int, count int) {
	submap := m[x]
	if submap == nil {
		submap = make(map[int]int, 0)
		m[x] = submap
	}
	m[x][y] = count
}

func isSet(m map[int]map[int]int, x int, y int) bool {
	// ignore intersections at the origin
	if x == 0 && y == 0 {
		return false
	}

	submap := m[x]
	if submap == nil {
		return false
	}
	return submap[y] != 0
}
