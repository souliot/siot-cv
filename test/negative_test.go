package test

import (
	"testing"
)

func TestNegativeGray(t *testing.T) {
	err := negativeGray("../data/am.jpg", "../data/am")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNegativeRGBA(t *testing.T) {
	err := negativeRGBA("../data/am.jpg", "../data/am")
	if err != nil {
		t.Fatal(err)
	}
}
