package mylib

import (
	"fmt"
	"regexp"
)

func init() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match) // true

	r := regexp.MustCompile("a([a-z]+)e") // 上記と意味は同じだけど, 何度も使うのであればこっちで書いたほうがいい
	ms := r.MatchString("apple")
	fmt.Println(ms) // true
	
	// s := "/view/test"をマッチさせたいとき
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs) // /view/test/
	
	// 一部取り出したいとき
	fss := r2.FindStringSubmatch("/view/test/")
	fmt.Println(fss, fss[0], fss[1], fss[2]) // [/view/test view test] , / view/test, view, test
}