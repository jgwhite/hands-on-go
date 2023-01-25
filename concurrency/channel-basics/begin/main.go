// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go sum(nums, ch)

	fmt.Printf("Result: %d\n", <-ch)

	ch2 := make(chan string)

	ch2 <- "James"
	ch2 <- "Toni"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
