package main

import "fmt"

func main() {
	// ARRAYS -----------------
	var fruitArray [2]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])

	// Declare & Assign Simultaneously
	vegArray := [2]string{"Carrot", "Tomato"}

	fmt.Println(vegArray)

	// SLICES --------------------
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1])
	fmt.Println(fruitSlice[2])
	// fmt.Println(fruitSlice[3]) // Index Out of Range Error

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2]) // [Begin: End+]
	fmt.Println(fruitSlice[1:3])
}
