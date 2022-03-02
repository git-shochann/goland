package main

import (
	"fmt"
	"strconv"
)

func plusone(nmstr string) (int, string) {
	res := ": "
	nm, err := strconv.Atoi(nmstr)
	if err != nil {
		res += err.Error() // なぜここでError()が使えるのか？ func Atoi(s string) (int, error)

	} else {
		nm++
		res += "おめでとうございます、" + nmstr
		res += "に1が加算されました。"
	}
	return nm, res
}

func main() {

	nm, res := plusone("53")
	fmt.Println(nm, res)

	nm, res = plusone("what")
	fmt.Println(nm, res)
}
