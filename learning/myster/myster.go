package main

import "fmt"

type person struct {
	name string
	age int
	smokes bool
	friend *person
}

func main() {
	myster := person{}
	fmt.Println("name:", myster.name)
	fmt.Println("age:", myster.age)
	fmt.Println("smokes", myster.smokes)
	fmt.Println("friend", myster.friend)
}

// name:
// age: 0
// smokes false
// friend <nil>