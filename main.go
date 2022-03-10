package main

import "fmt"

// this is execute file
func main() {
	slice := []string{"yua","reina","yuka"}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}