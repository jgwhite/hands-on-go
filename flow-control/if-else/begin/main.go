// flow-control/if-else/begin/main.go
package main

import (
	"fmt"
)

// parseOddsEvens returns two slices, one with the odd numbers and one with the even numbers
func parseOddsEvens(ints []int) ([]int, []int) {
	var odds []int
	var evens []int

	for _, i := range ints {
		if i%2 == 0 {
			evens = append(evens, i)
		} else {
			odds = append(odds, i)
		}
	}

	return odds, evens
}

func main() {
	// invoke and print the results of parseOddsEvens
	odds, evens := parseOddsEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(odds)
	fmt.Println(evens)
}
