package main

import (
	"errors"
	"fmt"
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

func TestFibonacci(t *testing.T) {
	c := require.New(t)

	result := fibonacciQuick(10)

	expected := 55

	c.Equal(expected, result)
}

func TestFact(t *testing.T) {
	c := require.New(t)

	testCases := []struct {
		number   int
		expected int
	}{
		{0, 1},
		{2, 2},
		{4, 24},
	}

	for i, e := range testCases {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			t.Parallel()
			out := fact(e.number)
			if out != e.expected {
				n := errors.New("error")
				c.NoError(n)
			}
		})
	}
}
