package main

import (
	"fmt"
	"strconv"
)

func calculate(input []string, verbose bool) (result float64, err error) {
	var stack []float64
	for _, token := range input {
		switch token {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*", "x":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/", ":":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "u":
			x := stack[len(stack)-1]
			y := stack[len(stack)-2]
			stack[len(stack)-1] = y
			stack[len(stack)-2] = x
		default:
			f, e := strconv.ParseFloat(token, 64)
			if e != nil {
				err = fmt.Errorf("%s: invalid!", token)
			}
			stack = append(stack, f)
		}
		if verbose {
			fmt.Println(stack)
		}
	}
	result = stack[len(stack)-1]
	return
}