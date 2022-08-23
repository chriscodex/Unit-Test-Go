package main

import (
	"testing"
)

func TestAddSuccessfully(t *testing.T) {
	result := Add(4, 5)

	expected := 9

	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
