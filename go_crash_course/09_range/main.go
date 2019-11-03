package main

import "fmt"

func main() {
	// RANGE WITH ARRAYS / SLICES -----------------
	ids := []int{22, 33, 44, 55, 66}

	// Loop through ids using RANGE
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Add all the ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// RANGE WITH MAPS ------------------------------
	emails := map[string]string{
		"Banna": "banna@gmail.com",
		"Kapil": "kapil@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s \n", k)
	}

	for _, v := range emails {
		fmt.Printf("%s, ", v)
	}
}
