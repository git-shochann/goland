バイト型

ネットワークとかファイルの読み込み系で使う？

[ASCII Code - The extended ASCII table](https://www.ascii-code.com/)

```go
func main() {

	b := []byte{72,73}

	fmt.Println(b) // => [72, 73] // アスキーコードのスライスが出力される

	fmt.Println(string(b)) // string型にキャストする => HI

}
```

---

**キャストするとは？**

変数のデータ型を異なるデータ型の変換すること

Go だと, 暗黙的な型変換は許されていないので, 下記のような異なる型の変数への代入はコンパイルエラーになる

```go
var i int = 500
var f float64 = i // Error! cannot use i (type int) as type float64
```

明示的な型キャストを使用する

```go
var i int = 100
var f float64 = float64(i)  // OK
```

---

基本的なエラーハンドリング

**コメント重要なので見るように！**

```go
func main() {
	file, err := os.Open("./hello.go") // ここのエラーを確認する -> Openの関数定義を見にいく
	if err != nil { // errがnilではない == エラーメッセージが返ってきている
	  fmt.Println(err) // open ./hello.go: no such file or directory
	  log.Fatalln("Error!") // 2021/12/22 19:17:52 Error!
    }
		defer file.close()
	  data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatalln("Error!")
    }
		fmt.Println(count, string(data))

		// オーバーライドで書き換える
		// 戻り値が1個の場合, 下記のように1つにまとめることが出来る
		if err = os.Chdir("test"); err := nil {
			log.Fatalln("Error!")
		}
}
```

※ panic と recover というのがあるけど、基本使わず、上記のエラーハンドリングを行うのがベスト

```go
package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect DB!")
}

func save() {
	defer func() {
		s:= recover() // ここでpanicをキャッチ出来る
		fmt.Println(s) // Unable to connect DB!
	}
	thirdPartyConnectDB()
}

func main() {
	save()
}
```

---

いろいろなコードをみて理解を深める

---

package

パッケージを読み込む

テストはアンダースコアで同階層に書く

main.go

lib -

      | - math.go

      | - math_test.go

```go
// math.go

package lib

type Person struct {
	name string
	age int
}

func Average(s []int) int {
	total := 0
	for _, i:= range s {
	  total += i
  }
	return int(total/len(s))
}
```

```go
// main.go

package main

import (
	"goland/lib"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(lib.Average(s))

	person := lib.Person{name: "Sho", age: 24}
	fmt.Println(person)
}
```

**ついでに testing も**

[testing](https://pkg.go.dev/testing)

`func TestXxx(*testing.T)`

```go
// math_test.go

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5}) // 実際に同じ呼び出す処理を書く
	if v != 3 { // 期待するのは3なので
		t.Error("Expected 3, got", t)
	}
}
```
