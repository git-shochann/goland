**ポインタ**

Go では, アドレスをポインタと呼び

ポインタが代入された変数をポインタ型変数と呼ぶ

ポインタを利用するとスコープを超えて値を更新出来る

ポインタ型変数の定義

変数のデータ型に`*` をつける

```go
func main() {
	name := "Yua"
  fmt.Println(&name) // 普通にnameのポインタを出力する
	var namePtr *string = &name // ポインタ型変数を宣言する
	fmt.Println(namePtr) // ポインタ型変数の出力を行う
}

// => 0x14000010230
// => 0x14000010230
```

ポインタ型変数を使って、直接値を更新することも可能

```go
var name := "yua"
fmt.Println(name) // => yua
fmt.Println(&name) // => メモリ上の位置(ポインタを参照)
var namePointer *string = &name // ポインタ型変数に代入
fmt.Println(namePointer) // ポインタを参照
*namePointer = "sho" // ポインタ型変数を更新
```

※1 ポインタを別の関数に引数として渡す場合, ※2 受け取る関数にはポインタ型の引数を用意する必要がある

```go
func main() {
	name := "Yua"
	changeName(&name) // ※1
	fmt.Println(name)
}

// ポインタ型の引数を用意する
func changeName(pointerArg *string) {
	*pointerArg = "Reina"
}
```

引数の復習

変数 totalScore を引数として渡すとき、その変数自体を渡している訳ではない。

実際には変数 totalScore の値がコピーされて、新しい変数に代入される → **値渡し**

```go
func main() {
	totalScore := 0
	fn(totalScore)
}

func fn(copyTotalScore int) {
}
```

引数の変数が別ポインタであることを調べる **後ほど?になったら実行してみる**

```go
package main

import "fmt"

func main() {
	totalScore :=  0
  fmt.Println(&totalScore)
	fmt.Println(&totalScore) // ポインタを参照
	test(totalScore)
}

func test(copyTotalScore int) {
	copyTotalScore += 10
	fmt.Println(copyTotalScore)
	fmt.Println(&copyTotalScore) // ポインタを参照
}
```

**ポインタを更新すればスコープを超えても元の値が更新される**

```go
package main

import "fmt"

func main() {
	totalScore := 10
	test(&totalScore)
	fmt.Println(totalScore) // => 110になる
}

func test(copyTotalScore *int) {
	*copyTotalScore += 100
}
```

**まとめ**

```go
package main

import "fmt"

func main() {
	a := 10
	b := 10

	// calculate関数を呼び出して、引数としてaそのままとbのポインタを渡してください。
	calculate(a, &b)

	fmt.Println("引数に整数を指定した場合：", a) // => 10
	fmt.Println("引数にポインタを指定した場合：", b) // => 11
}

func calculate(a int, bPtr *int) {
	// a, bPtrそれぞれに１を足す処理を記述してください
	a += 1
	*bPtr += 1
}
```

---

**ポインタ型使うとき...**

```go

func main() {
	totalScore := ask(1, "dog")
	totalScore += ask(2, "cat")
	totalScore += ask(3, "fish")

	fmt.Println("スコア", totalScore)
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
	    // 問題文の単語と入力値が一致するときの処理を追加してください
		fmt.Println("正解!")
		return 10


	} else {
	    // 問題文の単語と入力値が一致しないときの処理を追加してください
	    fmt.Println("不正解!")
		return 0

	}
}
```

**ポインタ型使わないとき...**

```go
package main

import "fmt"

func main() {
	totalScore := 0
	ask(1, "dog", &totalScore)
	ask(2, "cat", &totalScore)
	ask(3, "fish", &totalScore)

	fmt.Println("スコア", totalScore)
}

func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解!")
		// ポインタを使って加算してください
		*scorePtr += 10

	} else {
		fmt.Println("不正解!")
	}
}
```
