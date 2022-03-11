package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("ファイルの読み込みに失敗しました。", err)
		os.Exit(1)
	} else {
		cstr := string(content)
		lows := strings.Split(cstr, "\n")
		
		for i, v := range lows {
			if i > 0 && len(v) > 1 {
				data := strings.Split(v, ",")
				fmt.Printf("%sは%s円\n", data[0], data[1])
			} 
		}
	}
}