package main

import "fmt"

func main() {
	// My favorite color
	var favoriteColor = "Green"
	fmt.Println("My Favorite Color", favoriteColor)


	//My year of birth and age
	var birthYear = 2000
	var age = 2024 - birthYear

	fmt.Println("I was born in", birthYear, "I have", age)


	//Block 
	var (
		firstInitial = "M"
		surnameInitial = "S"
	)

	fmt.Println("Initials", firstInitial, surnameInitial)

	var ageInDays int
	ageInDays = age * 365

	fmt.Println("Age in days", ageInDays)


}