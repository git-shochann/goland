/*
This is the file to execute!
*/

package main

import (
	"fmt"
	"goland/lib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(lib.Average(s))

	person := lib.Person{name: "Sho", age: 24}
	fmt.Println(person)
}