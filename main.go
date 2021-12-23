/*
This is the file to execute!
*/

package main

import (
	"fmt"
	"goland/mylib"
)

func main() {
	fmt.Println("hi")
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	person := mylib.Person{Name: "Sho", Age: 24}
	fmt.Println(person)
}