// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func printVal() {
	fmt.Printf("global val = %v\n", val)
}

func setVal(v string) {
	val = v
}

func main() {
	// create a local variable "val" and assign it an int value
	val := 42

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printVal()

	// update the package-level variable "val" and print it again
	setVal("updated global")
	printVal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, &val = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Printf("local val = %v\n", val)

}
