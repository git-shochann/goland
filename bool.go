package main

import "fmt"

func main() {
	// var t, f bool = true, false
	t, f := true, false // short declaration
	fmt.Println(t, f)
	fmt.Printf("%T %v %t\n", t, t, t)
	// %T -> a Go-syntax representation of the type of the value => bool
	// %v -> the value in a default format
	// %t -> the word true or false
}