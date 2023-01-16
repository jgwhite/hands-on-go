// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "\033[1;31m%s\033[0m\n", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	if len(os.Args) <= 1 {
		panic("No file provided")
	}
	fn := os.Args[1]
	bs, err := os.ReadFile(fn)
	if err != nil {
		panic(fmt.Sprintf("Could not read file: %s", fn))
	}

	// convert the bytes to a string
	s := string(bs)
	fmt.Println(s)

	// initialize a map to store the counts
	c := map[string]int{
		"letters": 0,
		"symbols": 0,
		"numbers": 0,
		"words":   0,
	}

	// use the standard library utility package that can help us split the string into words
	w := strings.Split(s, " ")

	// capture the length of the words slice
	c["words"] = len(w)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, w := range w {
		for _, r := range w {
			if unicode.IsLetter(r) {
				c["letters"] += 1
			}
			if unicode.IsSymbol(r) {
				c["symbols"] += 1
			}
			if unicode.IsDigit(r) {
				c["numbers"] += 1
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(c)
}
