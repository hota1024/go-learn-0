package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func addSub(a int, b int) (int, int) {
	return add(a, b), sub(a, b)
}

func main() {
	fmt.Println(add(10, 24))
	fmt.Println(sub(10, 24))
	fmt.Println(addSub(10, 24))
}
