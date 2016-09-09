package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("addTwoInts: ", addTwoInts(1, 3))
}

func addTwoInts(a int, b int) int {
	return a + b
}
