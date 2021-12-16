**Slices(スライス)**

Go の  `Arrays(配列)`は固定長配列である一方、Go の  `Slices(スライス)`  は可変長配列であるのでより柔軟にデータを格納することが出来る

また、スライスは配列への**参照型**で値を持つため、配列から部分列を取り出す事**スライス操作**でも作成する事が可能

書き方例 3 選

```go
var testSlice [] string // 空っぽの変数宣言
fmt.Println(testSlice) // => []
```

```go
testSlice := [] int {2, 8} // int型のスライスを変数testSliceに代入
fmt.Println(testSlice) // => [2 8]
```

`変数名 := 配列[start:end]`

配列(またはスライス)から条件を指定して, スライスを作成する。

| 操作             | 意味                    |
| ---------------- | ----------------------- |
| Slice[start:end] | start から end - 1 まで |
| Slice[start:]    | start から最後尾まで    |
| Slice[:]         | 先頭から最後尾まで      |
| Slice[:end]      | 先頭から end - 1 まで   |

end - 1 とは？

以下参照 一旦実行してみれば分かる

```go
array := string{"Go", "Java", "TypeScript", "Ruby", "C#", "C++"}
slice := array[1:4]
fmt.Println(slice)

// => [Java, TypeScript, Ruby]
```

作成したスライスに値を入れると元々の配列も変更されるので注意！

```go
// 上記の続き

slice[0] = "たろちゃんです"
fmt.Println(array) // 元の配列を出力する
// => [Go たろちゃんです TypeScript Ruby C# C++]
```

**要素の追加**

`append()`

要素を末尾に追加し, 新しいスライスを返す

```go
slice := [] string{"Golang", "Java"} // スライスの作成
newSlice := append(slice, "Ruby") //sliceに"Ruby"を追加
```

スライスの長さと容量を求める

`len()` → 要素数

`cap()` → slice の生成元の要素数

```go
func main() {
	array := [...] string{"af1", "aj1", "am1"} // 配列の作成
	slice := array[0:1]
	fmt.Println(slice) // [af1]
	fmt.Println(len(slice)) // => 1
	fmt.Println(cap(slice) // => 3
}
```

スライス代入

スライスの型が一致している場合, スライスを代入することが出来る

```go
func main(){
     slice  := [] string{"Golang", "Java"}
     slice2 := []string //sliceと同型slice2を作成

     slice2 = slice //sliceをslice2に代入
     slice2[0] = "Ruby"

     fmt.Println(slice2) //=> [Golange Java]
}
```

スライスのゼロ値

nil スライスは 0 の長さと容量を持っており、何の元となる配列も持っていない。

```go
func main() {
    var slice []int
    fmt.Println(slice, len(slice), cap(slice)) //=> [] 0 0

    if slice == nil {
        fmt.Println("nilです!") //=> nil!
        //sliceの値がnilの場合にnil!を表示する。
    }
		// nilです!
}
```

**make()**

でもスライスを作ることが可能

第一引数が型, 第二引数が長さ, 第三引数が容量を示す

`make([]T, len, cap)`

```go
func main(){
	a:= make([]bool, 4, 4)
  fmt.Println(a) //=> 長さ5, 容量5ののbool型を作成する [false, false, false, false]
}
```

---

**連想配列(Maps というらしい...)**

特定のキーを用いて値を取り出せる配列

Ruby でいうハッシュ 難しい言葉だけど大したことまじでしてない

1, **Make を使って宣言する**

`make(map[キーの型]バリューの型, キャパシティの初期値)`

キャパシティ省略可

```go
mapExample := make(map[string]string, 2) //マップの宣言
fmt.Println(mapExample) // map[]

```

`make()`  によって作成された  `Maps(連想配列)`  に以下の様にキーとバリューを入れる事が出来る

`map[キーの値] = バリューの値`

```go
map["firstname"] = "Sho"
map["lastname"]  = "CEO"
```

実際に作ってみる

```go
func main() {
	mapExample := make(map[string]int, 2)
	fmt.Println(mapExample) // map[]
	map["age"] = 24
  map["favoriteNumber"] = 7
	fmt.Println(mapExaple) // map["age": 24, "favoriteNumber" : 7]
}
```

2, 初期値を指定して宣言

`var 変数名 map[キーの型]バリューの型 = map[key]value{key1: value1, key2: value2}`

```go
// 関数外では**Short variable declarations**不可
var mapExample = map[string] bool {"adult": true, "are you fine?" : true}

func main() {
	fmt.Println(mapExample)
}

// => map["adult": true, "are you fine?": true]
```

Maps(連想配列)のゼロ値

Maps(連想配列)の初期値を指定しない場合、変数は  `nil (nil マップ)`  に初期化される

`nil マップ`  は要素を格納することができず、

もし要素を格納したいのであれば、再度マップの初期化を行う必要がある。

→ 使うのかこれ？笑

```go
var map1 map[string]int

func main() {
    fmt.Println(map1) //=> map[] //nilマップ
}
```

**Maps(連想配列)への要素の挿入や更新**

要素の挿入

`map[キー] = バリュー`

```go
func main() {
	mapExample := make(map[string]int)

	mapExample["キー1"] = 100

	fmt.Println(mapExample) // map[キー1:100]

	mapExample["キー2"] = 500

	fmt.Println(mapExample) // map[キー1:100 キー2:500]
}
```

要素の取得

`map[キー]`

```go
func main() {

  mapEx := map[string] int { "age": 24, "favorite number" : 100 }

  fmt.Println(mapEx["age"]) // => 24

}
```

要素の削除

`delete(連想配列名, キー)`

```go
func main() {

  mapExample := map[string] string { "name" : "yua", "job" : "vocalist"}

  delete(mapExample, job) // 消して出力しないと！

	fmt.Println(mapExample) // => map["name" : "yua"]

}
```

---

Range

`Slices(スライス)`  や、`Maps(マップ)`  をひとつずつ反復処理するために for ループと一緒に使う(JavaScript の forEach みたいなやつ)

スライスを 1 つずつ処理する

`range`  は反復毎に 2 つの変数を返す

```go
var slice = []string{"Golang", "Ruby", "Javascript", "Python"} // スライスの作成

func main() {
	for index, value := range slice {
	fmt.Println(index, value)
}

// => 0 Golang
// => 1 Ruby
// => 2 Javascript
// => 3 Python
}
```

value の省略

`index, _` あ、この書き方は Udemy でやってた値がないとき`_`  へ代入することで省略するやり方だ。

```go
var slice = []string{"Golang", "Ruby", "Javascript", "Python"}

func main() {
	for index, _ := range slice {
	fmt.Println(index)
}
}
```

index の省略

```go
func main() {
	slice := []string {"Ramen", "Sushi", "Korean Food", "Gyudon"}
	for _, value := range slice {
	fmt.Println(value)
}
}
```

**反復処理を途中で止める、スキップする**

```go
var slice = []string{"Golang", "Ruby", "Javascript", "Python"}

func main() {
	for _, value := range slice {
	  if "Golang" == value {
	    fmt.Println("Golang Found!")
	    continue // 次の処理にいく
    }
		if "JavaScript" == value {
			fmt.Println("JavaScript Found!")
      break // ループ強制終了
		}
		fmt.Println(value")
  }
}

// Golang found
// Ruby is not Golang
// Javascript found!

// スライス内のPythonは呼ばれない //
```
