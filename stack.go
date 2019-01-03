package main

import (
	"fmt"
	"strconv"
)

func calculate(input []string, verbose bool) (result float64, err error) {
	if len(input) < 3 {
		err = fmt.Errorf("need 3 tokens minimum")
		return
	}
	var stack []float64
	for _, token := range input {
		switch token {
		case "+", "a", "p":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*", "x":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-", "s", "m":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/", ":", "d":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "u", "r":
			x := stack[len(stack)-1]
			y := stack[len(stack)-2]
			stack[len(stack)-1] = y
			stack[len(stack)-2] = x
		default:
			f, e := strconv.ParseFloat(token, 64)
			if e != nil {
				err = fmt.Errorf("token %q not valid", token)
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
