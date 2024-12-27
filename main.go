package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func orderString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if runes[j] < runes[minIndex] {
				minIndex = j
			}
		}
		runes[i], runes[minIndex] = runes[minIndex], runes[i]
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	var insertString string
	var orderFlag bool
	var inputString string

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	for _, arg := range args {
		if len(arg) > 9 && arg[:9] == "--insert=" {
			insertString = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertString = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			inputString = arg
		}
	}

	if insertString != "" {
		inputString += insertString
	}

	if orderFlag {
		inputString = orderString(inputString)
	}

	for _, r := range inputString {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
