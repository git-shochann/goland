package main

// import (
// 	"fmt"
// 	"image"
// 	"os"
// )

// func main() {
// 	reader, err := os.Open("matrinx.jpg") // return *File -> Read()メソッドを持っている
// 	if err != nil {
// 		fmt.Println("ファイルの読み取りエラーです", err)
// 		os.Exit(1) // プログラム終了
// 	}

// 	img, name, err := image.Decode(reader) // 引数にReaderインターフェースを取る -> Readインターフェースを満たしている構造体があればOK -> 戻り値の*File構造体はRead()メソッドを持っている
// 	if err != nil {
// 		fmt.Println("画像の変換エラーです")
// 		os.Exit(1)
// 	} else {
// 		fmt.Println(name, "形式のデータを取得しました")
// 	}

// 	defer reader.Close() // openしたらcloseする

// 	marks := []string("*", "+", "-")
// 	var marksStr string

// 	for y := img.Bounds()
// }