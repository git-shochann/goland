**インターフェイス**

インターフェイスは、構造体のように型が格納する値を定義するのではなく、

**何が実行されるかのメソッドの型を格納する**

```go
type インターフェース名 interface {
	メソッド()
}
```

interface の中にメソッドだけを書いて, そのメソッドをレシーバを使って具体的な処理コードを定義する

実際に書いてみる...

```go
package main

import "fmt"

// インターフェースの定義
type Scenario interface {
	dialog()
}

// 構造体の定義
type Character struct {
	name string
	rolename string
}

// メソッドを作る
func (c Character) dialog() {
	fmt.Println(c.name)
	fmt.Println(c.rolename)
}

// 実際にメソッドを呼び出す
func main() {
	var dragonball Scenario = Character{
	name: "Goku",
	rolename: "main",
  }
	dragonball.dialog()
}
```

空インターフェースというものもある

宣言した変数にはどんな型の値でも代入可能

```go
var obj interface{}

obj = 0123                                                     // int
obj = "String"                                                 // string
obj = []string{"Python", "Golang", "Ruby"}                     // slice
obj = func greetings(_ string) string { return "Hello World" } // function
```

---

型アサーションとは..？

インターフェースには型情報が欠落しているので

assertion → 断言, 断定

変数の型をチェックする機能

`i := a.(int)`

a が int 型か確認をして、OK なら変数に代入している

`value, ok := a.(int)`

また上記の様に書けば、1 番目の変数には**型アサーション成功時に実際の値が格納**され、

2 番目の変数には**型アサーションの成功の有無（true/false）**が格納される

```go
func main() {
	var a interface{} = 10
	i := a.(int)
	fmt.Println(i) // 10

	i, ok := a.(int)
	fmt.Println(i, ok) // 10, true

	f := a.(float32)
	fmt.Println(f) // panic: interface conversion: interface {} is int, not float32
  // 成功したかの有無を確かめるokが存在しないのでエラーが発生してしまう

	f, ok := a.(float32)
	fmt.Println(f, ok) // 0, false -> エラーは返さない
}
```

この型アサーションは**型に応じた分岐処理**の実現が可能。

事前に型情報が不明なものに、実行時に型に応じた処理をさせることができる。

**型 Switch**

データ型の判定を行う書き方

`switch`  のあとに, 型アサーション  `v := x.(type)`  と書き、`case`  に型を指定

```go
func main() {
	var a interface{} = 10
	switch variable := a.(type) {
		case int:
			fmt.Println(variable)
		case string:
			fmt.Println(variable)
		default:
			fmt.Println("default")
  }
}
```

```go
func main() {
    var a interface{} = 10
    switch a.(type) {
    case int:
        fmt.Println("This is int.")
    case float32:
        fmt.Println("This is float.")
    case string:
        fmt.Println("This is string.")
    default:
        fmt.Println("This is unknown type.")
    }
}
```

**補足**

switch 文初めに 変数に代入している意味分からない → とりあえず理解した！

ここの変数  `v`  はそれぞれ `型int`  または `型string`  であり、 `**i`  によって保持される値を保持する\*\*

```go

// main関数内で引数をもった状態で呼び出し -> 下記のiが対応引数となる(iはinterface型)

func do(i interface{}) {
	switch v := i.(type) {
	  case int:
		  fmt.Printf("値は%d, 型は%T", v, v)
		case string:
			fmt.Printf("値は%d, 型は%T", v, v)
		default:
			fmt.Println("デフォルトが呼び出されました")
  }
}

func main() {
	do(24)
	do("Hello World")
}
```

これわかりやすい。

[Go 言語 - 空インターフェースと型アサーション - 覚えたら書く](https://blog.y-yuki.net/entry/2017/05/08/000000)

---

**とりあえず書く関数, メソッド, 構造体, インターフェース等の流れは実践の部分でやってみる！**

---
