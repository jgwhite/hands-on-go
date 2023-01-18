// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Printf("printString(\"%s\") # => %s\n", s, s) }

func printInt(i int) { fmt.Printf("printInt(%d) # => %d\n", i, i) }

func printBool(b bool) { fmt.Printf("printBool(%v) # => %v\n", b, b) }

type printable interface {
	~string | ~int | ~bool
}

func printAny[T printable](v T) {
	fmt.Printf("printAny(%v) # => %v\n", v, v)
}

// Part 2 sum function refactoring

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

type numeric interface {
	constraints.Integer | constraints.Float
}

func sumAny[T numeric](ns ...T) T {
	var result T

	for _, n := range ns {
		result += n
	}

	return result
}

func main() {
	// call non-generic print functions
	printString("Hello")
	printInt(42)
	printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	fmt.Println("sum() # => ", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println("sumAny() # => ", sumAny(1, 2, 3))
}
