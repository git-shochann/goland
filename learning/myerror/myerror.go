package main

import "fmt"

type myerror struct {
	msg string
}

// myerrorインスタンスのメソッドを作成する
func (e *myerror) Error() string {
	return e.msg
}

func read(word string) (string, error) { // 戻り値はstring型とerror型
	if len(word) > 0 { // 文字列の長さもlenで求められる
		return word + "と言いましたね", nil // 条件クリアでstringとnilを返す
	}
	return "", &myerror{"読み取れませんでした"} // 条件失敗で空の文字列とmyerrorインスタンスを返す
}

func main() {
	words := []string{"こんにちは", "", "こんばんは"} // stringのスライスを作成
	for i:=0; i < len(words); i ++ { // 左辺が右辺よりも小さいかどうか
		result, err := read(words[i])
		if err != nil { // errに値が入っていたら
			// fmt.Printf("値は【%s】, 型は%T\n", err, err) // 値は【読み取れませんでした】, 型は*main.myerror
			fmt.Println(err.Error()) // エラーインスタンスのメソッドError()を発動
		} else {
			fmt.Println(result)
		}
	}
}