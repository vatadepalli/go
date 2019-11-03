package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(greeting("Banna"))
	fmt.Println(getSum(3, 9))
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func greeting(name string) string {
	return "Hello " + name + "!"
}
