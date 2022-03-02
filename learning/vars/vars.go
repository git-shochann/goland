package main

import "fmt"

type mynumber struct {
	value int
}

func main() {

	var something int
	var someone string
	var somenumber mynumber
	var somearr [5]int


	fmt.Println("somearr:", somearr)
	fmt.Println("something:", something)
	fmt.Println("somenumber", somenumber)
	if someone == "" {
		fmt.Println("空の文字列です")
	} else {
		fmt.Println(something, someone)
	}
}

// something: 0
// 空の文字列です