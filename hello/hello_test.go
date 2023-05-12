package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9))
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3))
}
