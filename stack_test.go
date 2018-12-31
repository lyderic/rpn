package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	inputs := [][]string{
		[]string{"56", "9", "x"},
		[]string{"56", "9", "*", "88", "-", "5.005", "/"},
		[]string{"56", "9", "^"},
		[]string{"2", "0", "+"},
		[]string{"123", "0", ":"},
		[]string{"3", "4", "2", "+", "x"},
		[]string{"45", "90", "4", "x", "u", ":"},
	}
	for _, input := range inputs {
		display(input)
	}
}
