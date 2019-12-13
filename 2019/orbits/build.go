package orbits

import (
	"fmt"
	"strings"
)

// Orbits TODO godoc
type Orbits struct {
	count  int
	orbits map[string]*Orbit
}

// Orbit TODO godoc
type Orbit struct {
	source string
	orbits *Orbit
}

// Build TODO godoc
func Build(source string) Orbits {
	o := Orbits{}

	o.orbits = make(map[string]*Orbit, 0)

	o.parse(source)

	return o
}

// Path TODO godoc
func (o Orbits) Path(source string) []string {
	so := o.orbits[source]
	return so.Path()
}

// Count TODO godoc
func (o *Orbits) Count() int {
	if o.count == 0 {
		o.count = o.countOrbits()
	}
	return o.count
}

func (o Orbits) countOrbits() int {
	total := 0

	for k, v := range o.orbits {
		n := v.countOrbits()
		total += n
		fmt.Printf("%s has %d orbits. Total is now %d.\n", k, n, total)
	}
	return total
}

func (o *Orbits) parse(source string) {
	paths := strings.Fields(source)

	for _, path := range paths {
		parts := strings.Split(path, ")")
		target := parts[0]
		source := parts[1]

		o.addOrbit(source, target)
	}
}

func (o *Orbits) addOrbit(source string, target string) {
	so := o.getOrbit(source)
	to := o.getOrbit(target)

	so.orbits = to
	fmt.Println(source + " orbits " + target)
}

func (o *Orbits) getOrbit(name string) *Orbit {
	orb, ok := o.orbits[name]
	if ok {
		return orb
	}

	orb = &Orbit{
		source: name,
	}

	o.orbits[name] = orb

	return orb
}

func (o Orbit) countOrbits() int {

	return o.countOrbitsRecursive(0)
}

func (o Orbit) countOrbitsRecursive(n int) int {
	if o.orbits != nil {
		return o.orbits.countOrbitsRecursive(n + 1)
	}

	return n
}

// Path TODO godoc
func (o Orbit) Path() []string {
	return o.pathRecursive([]string{})
}

func (o Orbit) pathRecursive(p []string) []string {
	if o.orbits != nil {
		p = append(p, o.source)
		return o.orbits.pathRecursive(p)
	}

	return p
}

func shortestPath(path1 []string, path2 []string) []string {

	path := make([]string, 0)

	for _, p1 := range path1 {
		path = append(path, p1)
		pp := pathTo(path2, p1)
		if len(pp) > 0 {
			// compute shortest path
			return append(path, pp...)
		}
	}
	return path
}

func pathTo(p []string, v string) []string {
	path := make([]string, 0)
	for _, e := range p {
		if e == v {
			return path
		}

		path = append(path, e)
	}
	return []string{}
}
