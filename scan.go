package main

import "fmt"

func main() {
	name := ""

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %s!", name)
}
