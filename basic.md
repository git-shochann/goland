// 変数宣言 var → variable

var 変数名 データ型
var message string

// 再代入

message = "Hello, Go"

// データ型の省略

var age = 24

// さらに簡単にする (var, データ型の省略)

age := 24

// println で 2 つ値を出力する

name := "tarochan"
age := 3
println(name, age)

/_
使っていない変数の存在はバグ（不具合）の原因となることが多いため、
Go ではエラーを発生させて、バグを未然に防ぐ設計
_/

// 自己代入

n := 100
n = n + 20

// 省略形

n += 20 // n + 20 をした結果を n に代入
n %= 30 // n % 30 をした結果を n に代入

// 便利な書き方

n++ // 変数に 1 を足す
n-- // 変数から 1 を引く

// 🐾 大小を比べる

x < y // x が y より小さい時成り立つ
x <= y // x が y より小さいまたは等しい時成り立つ

// 🐾 等しいかを比べる

== // 等しい
!= // 等しくないとき true

// 全ての条件式が true の場合のみ全体として true

true && true
-> true

// どれかが true の場合 true を返す

true || false || true
-> true

/_
switch 文
条件分岐が多い時は、if 文を使うよりもシンプルに書ける場合がある
どの case とも一致しない場合, defalut が実行される
_/

switch (条件式) {
case 値 1 :

    case 値2, 値3 :

default:
}

// fmt.Printf
fmt.Printf(書式, 出力に用いる値);

weather := "Summer"
fmt.Printf("今日の天気は%s です", weather)

month := 8
day := 25
fmt.Printf("今日は、%d 月%d 日です", month, day)

// 改行されない
func main() {
fmt.Printf("Hello %s さん", "Sho")
fmt.Printf("Hello %s さん", "Taro")
}
// => Hello Sho さん Hello Taro さん

// 改行 \n を使う
fmt.Printf("Hello %s さん"\n, "Sho")
fmt.Printf("Hello %s さん", "Taro")

/\*

for 文 一定の処理を自動で繰り返し行う処理

for ① 変数の初期化; ② 繰り返しの条件, ④ 変数の更新 {
③ fmt.Println(i);
}

言葉でまとめると...

1, 変数宣言
2, 条件
3, 繰り返す処理
4, 変数の更新

- for 文の変数宣言で var は使えない \*

\*/

func main() {
for i:=1; i <= 3; i++ {
fmt.Println(Hello)
fmt.Println(i)
}
}

// => Hello 1
// => Hello 2
// => Hello 3

/_
ランダムな数を扱う math/rand
ランダムの数を日本語で乱数という
_/

package main

import "fmt"
import "math/rand"

func main() {

    for i := 1; i <= 5; i++ {
    // 0 ~ 9の乱数を5回表示させる。
        fmt.Println(rand.Intn(10))
    }

}

// ただ上記では、毎回実行するたび同じ乱数が表示されてしまう

// 比較演算子の復習 たまに分からなくなる

i <= 5
i => 5
i < 5
i > 5
