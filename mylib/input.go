package mylib

import (
	"bufio"
	"fmt"
	"os"
)

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}