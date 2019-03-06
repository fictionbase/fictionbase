package fictionbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCw(t *testing.T) {
	c := NewCw()
	assert.Nil(t, nil)
	assert.NotNil(t, c)
}
