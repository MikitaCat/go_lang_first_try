package main

import "fmt"

func main() {
	var myName = "Mikita"
	fmt.Println("My name is", myName)

	var oneMoreName string = "Ursula"
	fmt.Println("Other name is", oneMoreName)

	userName := "Admin"
	fmt.Println("Other name is", userName)

	var withoutValueInt int
	fmt.Println("Variable sum,", withoutValueInt)

	a, b := 4, 5
	fmt.Println("Number a=", a, "Number b=", b)

	// Trying to rewrite a variable
	a, h := 9, 8
	fmt.Println("Number a=", a, "Number b=", b, "Number h=", h)

	sum := a + b
	fmt.Println("Sum of a and b", sum)

}