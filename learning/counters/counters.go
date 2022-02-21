package main

import "fmt"

func counter() func() {
	fmt.Println("initialize...")
	ctr := 0
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func main() {
	countfunc := counter()
	fmt.Printf("%T", countfunc) // func()
	// countfunc()
	// countfunc()
	// countfunc()
}
