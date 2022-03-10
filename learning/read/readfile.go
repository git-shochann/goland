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
		fmt.Println("ファイルの読み込みに失敗しました。", err) // ファイルの読み込みに失敗しました。 open daata.txt: no such file or directory
		os.Exit(1)
	} else {
		cstr := string(content) // -> string配列
		lows := strings.Split(cstr, ",") // 文字列を区切って更に配列を作成する
		
		for i, v := range lows {
			if i > 0 && len(v) > 1 {
				data := strings.Split(v, ",")
				fmt.Printf("%sは%s円\n", data[0], data[1])
			} 
		} // 配列をrangeで回す 
	}
}

// まずはfor range
// 流れをUdemy