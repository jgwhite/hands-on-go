// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	var i any = "hello"
	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	}
}
