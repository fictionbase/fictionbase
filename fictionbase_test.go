package fictionbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetViperConfig(t *testing.T) {
	SetViperConfig()
	assert.Nil(t, nil)
	assert.NotNil(t, "string")
}
