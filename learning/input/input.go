package main

import "fmt"

type input interface {
	read() string
}

type keyboard struct {}

type file struct {}

type mic struct {}

type senser struct {}

func (k keyboard) read() string {
	return "何かキーを押してください"
}

func (f file) read() string {
	return "ファイルパスを指定してください"
}

func (m mic) read() string {
	return "ボタンを押して話してください"
}

func (s senser) read() string {
	return "計測を開始してください"
}

func main() {
	// スライスの作成？
	method := []input{
		keyboard{}, file{}, mic{}, senser{},
	}
	// iが4よりも小さいところまで繰り返す
	for i:=0; i < len(method); i++ {
		fmt.Println(method[i].read())
		// 1回目: keyboardインスタンス.read()
		// "何かキーを押してください"
	}
}