package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from function")
}

func greet(name string) {
	fmt.Println("Hello,", name, "nice to see you")
}

func sum(a, b int) int {
	return a + b 
}

func main() {
	sayHello()
	greet("Mikita")
	greet("Nokt")

	var number = sum(10, 12)
	fmt.Println("Number as a func result", number)
}