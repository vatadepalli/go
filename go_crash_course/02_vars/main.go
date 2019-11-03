package main

import "fmt"

func main() {
	// Data Types
	/*
		string
		bool
		int
		int 	int8 	int16 	int32 	int64
		uint 	uint8 	uint16 	uint32 	uint64 	uintptr
		byte - alias for uint8
		rune - alias for int32
		float32 	float64 (default)
		complex64 	complex128
	*/

	// CRETAING VARIABLES

	// USING VAR
	var name = "Banna"

	// var age = 37
	var age int32 = 37
	// When you delclare a variable in Go you have to use it.
	var size = 2.3

	fmt.Println(name)
	fmt.Printf("%T \n", name)

	fmt.Println(age)
	fmt.Printf("%T \n", age)

	fmt.Println(size)
	fmt.Printf("%T \n", size)

	// USING CONST
	const isCool = true
	fmt.Printf("%T \n", isCool)
	// isCool = false	// Throws an error - Can't assign to const

	// SHORT HAND - Only possible inside functions
	branch := "DAC"
	fmt.Println(branch)

	// MISC
	name, email := "kapil", "katsuper@hotmail.com"
	fmt.Println(name, email)
}
