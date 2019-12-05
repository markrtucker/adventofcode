package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindIntersection(t *testing.T) {

	d := FindIntersection("U3,R4,D2", "L1,U2,R14,D1")

	assert.Equal(t, 2, d)
}

// func Test_FindIntersection2(t *testing.T) {

// 	d := impl.FindIntersection("R8,U5,L5,D3", "U7,R6,D4,L4")

// 	assert.Equal(t, 3, d)
// }
