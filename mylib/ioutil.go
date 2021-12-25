package mylib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func Ioutil() {
	
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content)) // 戻り値は生のデータのbyte(ASCIIコード)のスライスなので, 見えるようにstringにキャストする
	
	content = []byte("hello world")
	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewBuffer([]byte("abc")) // バッファを作成する
	content, _ = ioutil.ReadAll((r)) // バイトスライスを読み込む
	fmt.Println(string(content))
	
}