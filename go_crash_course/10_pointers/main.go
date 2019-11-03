package main

import "fmt"

func main() {
	/*
		POINTERS
		point to memory locaions
	*/

	a := 5
	b := &a

	fmt.Println(a)
	fmt.Printf("%T \n", a)

	fmt.Println(b)
	fmt.Printf("%T \n", b)
}
