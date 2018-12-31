package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fix int // fixed decimals
	var verbose bool
	flag.IntVar(&fix, "f", 2, "fix")
	flag.BoolVar(&verbose, "v", false, "verbose stack")
	flag.Parse()
	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input := strings.Fields(string(bytes))
		display(input, fix, verbose)
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			print("rpn> ")
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			switch line[0] {
			case 'q':
				os.Exit(0)
			case 'v':
				verbose = !verbose
				fmt.Println("verbose:", verbose)
				continue
			case 'f':
				if len(line) == 1 {
					fmt.Println("fix:", fix)
				} else {
					decimals := string(line[1:])
					var err error
					fix, err = strconv.Atoi(decimals)
					if err != nil {
						fmt.Println("cannot set fix to", fix)
					}
				}
				continue
			default:
				input := strings.Fields(line)
				display(input, fix, verbose)
			}
		}
	}
}

func display(input []string, fix int, verbose bool) {
	result, err := calculate(input, verbose)
	if err != nil {
		fmt.Printf("Input: '%#v' failed! %v\n", input, err)
	} else {
		format := fmt.Sprintf("%%.%df\n", fix)
		fmt.Printf(format, result)
	}
}
