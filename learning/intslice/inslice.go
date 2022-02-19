package main

import "fmt"

func main() {
	intvals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("this number is %d\n", len(intvals))
	fmt.Printf("last element is %d \n", intvals[len(intvals)-1])
}
