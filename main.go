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

const (
	VERSION = "0.0.3"
)

func main() {
	var fix int // fixed decimals
	var verbose bool
	flag.IntVar(&fix, "f", 2, "fix `decimals`")
	flag.BoolVar(&verbose, "v", false, "verbose stack")
	flag.Usage = usage
	flag.Parse()
	fifo, _ := os.Stdin.Stat()
	var input []string
	if (fifo.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input = strings.Fields(string(bytes))
		display(input, fix, verbose)
		return
	}
	if len(flag.Args()) > 0 {
		input = flag.Args()
		display(input, fix, verbose)
		return
	}
	version()
	reader := bufio.NewReader(os.Stdin)
	for {
		print("rpn> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		switch line[0] {
		case 'q', 'x':
			os.Exit(0)
		case 'v':
			verbose = !verbose
			fmt.Println("verbose:", verbose)
			continue
		case 'h':
			commands()
			continue
		case 'V':
			fmt.Println("version:", VERSION)
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

func display(input []string, fix int, verbose bool) {
	result, err := calculate(input, verbose)
	if err != nil {
		fmt.Printf("Invalid input! %v\n", err)
	} else {
		format := fmt.Sprintf("%%.%df\n", fix)
		fmt.Printf(format, result)
	}
}

func version() {
	fmt.Printf("rpn v.%s - (c) Lyderic Landry, London 2019\n", VERSION)
}

func commands() {
	fmt.Println(`Commands:
  +, a, p     addition
  -, s, m     substraction
  *, x        multiplication
  /, :, d     division
  u, r        swap X <-> Y`)
}

func usage() {
	version()
	fmt.Println("Usage: rpn <options> <expression>")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println(`Examples:
  $ echo '3 3 x' | rpn
  $ echo '45 90 4.004 * u /' | rpn -f 3
  $ rpn 1 2 3 4 5 x x x x`)
}
