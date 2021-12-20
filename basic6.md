**並行処理**

`goroutine`

Go のプログラムが並行で実行されるもの

**goroutine の起動**

関数またはメソッドの呼び出し前に go をつけると、異なる`goroutine`で関数を実行する

新しい`goroutine`は並行に動作するので、前の関数の実行終了を待つことなく、次の`go`のあとに記述されているプログラムを実行する

```go
go 関数名(引数, ...)
```

**goroutine の終了条件**

- 関数の処理が終わる
- `return`で抜ける
- runtime パッケージを導入して, `runtime.Goexit()`を実行する

存在している goroutine を取得してみる

```go
import (
	"log"
	"runtime"
)

func main() {
	log.Println(runtime.NumGoroutine())
}
```

goroutine を使用せず、実行結果を見てみる

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	process(2, "A")
	process(2, "B")
	fmt.Println("Finish!")
}

func process(num int, str string) {
	for i := 0 ; i <= num ; i++ {
	  time.Sleep(1 * time.Second) // 処理を待機する
		fmt.Println(i, str)
    }
}

// Start!
// 0 A
// 1 A
// 2 A
// 0 B
// 1 B
// 2 B
// Finish!
```

goroutine で書いてみると...?

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	go process(2, "A")
	go process(2, "B")
	fmt.Println("Finish!")
}

func process(num int, str string) {
	for i := 0 ; i <= num ; i++ {
	  time.Sleep(1 * time.Second) // 処理を待機する
		fmt.Println(i, str)
    }
}

// =>
// Start!
// Finish!
```

上記だと main 関数内においての, goroutine の終了はまたず main 関数が終了してしまう

→ `関数process()`の実行結果を得るためには, **main が goroutine が終了するまで, 待たなくてはいけない**

**Channel**を使うことで, 解決！

Channel(チャネル)

goroutine 間(main を含む)で連携するには, channel と呼ばれる機能を利用する

値の交換, 同期という通信機能を備えていて, 2 つの goroutine が予期しない状態とならないことを保証する仕組み

**Channel を生成する**

組み込み関数 make()を使って, 宣言する

```go
ch := make(chan 型)
ch := make(chan 型, バッファサイズ)
```

バッファサイズとは

一時的に記憶可能な場所 → バッファ

バッファ可能な容量で, このサイズが送信データの上限となる

**チャネルでの値の送信, 受信方法**

```go
ch <- 構文 // 構文をchに送信する

args := <- ch // chから受信した変数をargsに割り当てる
```

**実際に上のを踏まえて, 書き直してみる。**

```go
func main() {
	channel := make(chan string) // channelの作成

  go func() { channel <- "Good Evening!" }() // 作成したchannelに, 値(str)を作成

	message := <- channel // channelから値を受信

	fmt.Println(message)
}

// => Good Evening!
```

かなり分かりやすい

[Go の goroutine, channel をちょっと攻略！ - Qiita](https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e)

追記 : goroutine の同期

**送る側と受ける側が準備できるまで, 送受信はブロックされる**

このため、同期処理的なものを書かなくていい。

なのでその特性を利用して...

---

EX1

```go
func main() {

  ch := make(chan bool) // bool型のchannelを作成

	go func() {
		fmt.Println("Hello!")
		ch <- true // 値を送信する, 値はなんでもOK(bool型であれば)
	}()

	<- ch // channelの型であるboolの値を受け取るまでの完了待ち

}
```

main 内で`ch`の型である bool の値を受け取るまで完了待ち

---

EX2

**コメントしっかり読むこと**

```go
func hello(done) {
	fmt.Println("Hello!")
	done <- true // ③ channelに送信する
}

func main() {
	done := make(chan bool) // ① channelの作成
	go hello(done) // ② 生成したchannelを関数に渡す
  <- done // ④ ③が終わる前にここの処理が来るけど, bool型の値doneを受信するまで待機

	fmt.Println("channelからの返信をもらったので, hello関数を終了します。")
}
```

---

**上記を学習したので, ちゃんと動くように書いてみる**

EX3

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool) // チャネルの作成
	ch2 := make(chan bool) // チャネルの作成

	fmt.Println("Start!")

	go func() {
		process(2, "A")
		ch1 <- true
	}

	go func() {
		process(2, "B")
		ch2 <- true
  }

	fmt.Println("Finish!")
}

<- ch1 // ch1の受信を待つ
<- ch2 // ch2の受信を待つ

func process(num int, str string) {
	for i := 0 ; i <= num ; i++ {
	  time.Sleep(1 * time.Second)
		fmt.Println(i, str)
    }
}

// => Try it out!

```

`<-ch1`や`<-ch2`が呼ばれると、bool 型の要素を受信するまで完了待ちをする

---

基礎一通り触った！

- 所要時間 : 時間は置いておいて触った日数で大体 6 日 😘

次は Udemy で復習も踏まえて(ここを見に来る), 実践をする
