package main

import "fmt"

// Cat structure.
type Cat struct{}

// Meow meow.
func (c Cat) Meow() {
	fmt.Println("にゃーん")
}

func main() {
	cat := new(Cat)
	cat.Meow()
}
