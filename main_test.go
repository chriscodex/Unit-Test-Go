package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccessfully(t *testing.T) {
	result := Add(4, 5)

	expected := 9

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

// Clean test
func TestAddSuccessfullyTestify(t *testing.T) {
	c := require.New(t)

	result := Add(100, 150)

	expect := 250

	c.Equal(expect, result)
}
