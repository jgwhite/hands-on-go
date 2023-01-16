// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) speak() { fmt.Printf("%s speaking...\n", p.name) }
func (p person) walk()  { fmt.Printf("%s walking...\n", p.name) }

func (p person) String() string { return fmt.Sprintf("Hello, it’s %s", p.name) }

// define a dog type that can walk but not speak
type dog struct{}

func (d dog) walk() { fmt.Println("Dog walking...") }

func main() {
	// invoke with a person
	doHumanThings(person{name: "Johnny"})

	// can we invoke with a dog?
	// doHumanThings(dog{})

	fmt.Println(person{name: "Johnny"})
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
