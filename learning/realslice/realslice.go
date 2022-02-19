package main

import "fmt"

func main() {
	intarray := [5]int{98, 125, 232, 147, 486}
	slice0 := intarray[4:]
	fmt.Println(slice0)
}
