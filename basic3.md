なんだこれメモ, Progate で足りていなかった補足(Qiita, Youtube レッスンも含む)

`var c, python, java bool`

ただの変数定義の省略形で中身は何もない

```go
-> var c bool
-> var python int
-> var java string

fmt.Println(c, python, java)

// => false, false, false
```

※ 変数に初期値を与えないと、Zero values が与えられるが、それは型によって変わる( 上記参照 )

---

関数外では**Short variable declarations**は使用不可

---

定数`const`は  `:=`  を使って宣言出来ない

---

if ステートメントは条件の前に, 評価するための簡単なステートメントを書くことが出来るけど,

そこでの変数は if のスコープ内で有効

---

defer ステートメント

`defer`  へ渡した関数の引数は、すぐに評価されるけど, その関数自体は呼び出し元の関数が return するまで実行されない。

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// => "hello"
// => "world"
```

複数ある時は逆の順で実行される

[https://go-tour-jp.appspot.com/flowcontrol/13](https://go-tour-jp.appspot.com/flowcontrol/13)

---

**Struct (構造体)**

異なるデータ型の変数を 1 つにまとめる

中身は**フィールド**という

```go
type Student stract {
	name string
	math, english int
}

func main() {
	var s Student // 構造体の初期化
	s.name = "Sho"
	s.math = 98
	s.english = 100
}

// 構造体が代入された変数.フィールド名

fmt.Println(s)

// => { Sho, 98, 100 }
```

**構造体を初期化する方法は 2 通り**

- 変数定義後にフィールドを設定する方法 (上記)

- `{}`  で順番にフィールドの値を渡す方法

この場合であればフィールド通りで書かないといけないので、

そんな時は初期化時, フィールド名を書く { age: "24", firstName: "Tsuboya" }

```go
type Person struct {
   firstName string
   age int
}

func main(){
  sho := Person{"Bob", 24}
  fmt.Println(sho.firstName, sho.age) // フィールドにアクセスする
}
```

---

**メソッド**

Go ではクラスの概念がないので、別言語とは異なる

**構造体などの特定の型に関連づけられた関数？**

要するにメソッドは構造体を使用して使う

```go
type Student stract {
	name string
	math, english int
}

// メソッドの作り方

func (s Student) avg() {
	fmt.Println(s.name, (s.math + s.english) / 2 )
}

func main() {
	a001 := Student("Yua", 80, 70)
	a001.avg()
}
```

```go
func (構造体から初期化した変数名 構造体名)

// 上記()の部分をレシーバーという
```

まとめ

```go
// 構造体の定義
type Person struct {
	firstName string
	age int
}

// メソッドの定義
func (p Person) intro(greeting string) string {
	return greeting " I am " + p.firstName
}

func main() {
	Sho := Person{"Sho", 24}
	fmt.Println(Sho.intro("Hello")) // => Hello I am Sho
}
```

---

**継承に似た, 構造体の埋め込みがある**

```go
package main

import "fmt"

type Person struct {
	firstName string
}

func (a Person) name() string {
	return a.firstName
}

// 構造体の埋め込み
type User struct {
	Person
}

func (a User) name() string {
	return a.firstName
}

func main() {
	Sho := Person{"Tsuboya"}
	Yua := User{}
	Yua.firstName = "sho"
	fmt.Println(Sho.name())
	fmt.Println(Yua.name())
}
```

---

データ型

Arrays(配列)

**固定長配列**なので,最初に宣言した配列のサイズを変更は後から出来ない

`変数 := [要素数] データ型 {データ1, データ2}`

基本的な書き方

```go
a := [3] string {"Sho", "Taro", "Yua"}
a[1] = "Reina" // 配列のデータを変更
```

配列の要素の省略をする

```go
a := [...] string {"Sho", "Taro", "Yua"}
```

二次元配列

```go
a := [2][2] string {{"Sho", "Yua"}, {"Sho", "Reina"}}

// データの取得
fmt.Println(a[0][1])

// => Yua
```

---
