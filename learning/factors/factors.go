package main

import "fmt"

func getFactors(thenum int) {
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0:
		case n == 1:
		case thenum%n == 0:
			thenum /= n
			fmt.Printf("%dで割ると, %dです\n", n, thenum)
		default:
		}
	}
	fmt.Println()
}

func main() {
	getFactors(2000)
	getFactors(3000)
}
