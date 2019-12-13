package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MeetsRules(t *testing.T) {

	// assert.True(t, MeetsRules(122456))
	// assert.True(t, MeetsRules(111111))

	assert.False(t, MeetsRules(122256))
	// assert.False(t, MeetsRules(123456))
	// assert.False(t, MeetsRules(123454))

	// assert.False(t, MeetsRules(123454))
	// assert.False(t, MeetsRules(123454))
	// assert.False(t, MeetsRules(123454))
}
