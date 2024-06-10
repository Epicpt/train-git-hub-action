package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	get := Hello()
	want := "Hello, world!"
	assert.Equal(t, get, want)
}
