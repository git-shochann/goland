package main

import "fmt"

// func main() {
// 	x := mylib.Input("Type a number!")
// 	fmt.Print(x + "月は")
// 	switch n, err := strconv.Atoi(x); n {
// 	case 0: // p.72
// 		fmt.Println("整数値が得られません")
// 		fmt.Println(err)
// 	case 1, 2, 12:
// 		fmt.Println("冬です")
// 	case 3,4,5:
// 		fmt.Println("春です")
// 	case 6,7,8:
// 		fmt.Println("夏です")
// 	case 9,10,11:
// 		fmt.Println("秋です")
// 	default:
// 		fmt.Println("月の値ではないですよ？")
// 	}
// }

// func main() {
// 	x := mylib.Input("Type a number")
// 	n, err := strconv.Atoi(x)
// 	if err == nil {
// 	} else {
// 		return // main関数を終了するという意味
// 	}
// 	switch {
// 		case n%2 == 0:
// 			fmt.Println("偶数です")
// 		case n%2 == 1:
// 			fmt.Println("奇数です")
// 	}
// }

// func main() {
// 	// スライスの作成
// 	a := []int{10,20,30}
// 	// 関数を呼び出す
// 	a = push(a, 1000)
// 	// 変数aを出力する
// 	fmt.Println(a)
// }

// // スライスの末尾に加える関数
// func push(a []int, v int) []int {
// 	return append(a, v)
// }

// func main() {
// 	a := []string{"apple","google","meta"}
// 	b, c := push(a,"amazon")
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// // string型配列とstringの値を受け取り、配列の末尾に追加する関数を作成する
// // 戻り値は加えたスライスとスライスの要素数を返す
// func push(a []string, v string)([]string, int) {
// 	return append(a, v), len(a)
// }

// 無名関数
// func main() {
// 	//
// 	f := func(a []string)([]string, string){
// 		return a[1:], a[0]
// 	}

// 	// スライスを作成
// 	m := []string{"one", "two", "three"}

// 	s := ""

// 	fmt.Println(m)

// 	// 繰り返す
// 	for len(m) > 0 {
// 		m,s = f(m)
// 		fmt.Println(s + "->" , m)
// 	}

// }

// func main() {

// 	// 第一引数にstringのスライス, 第二引数にf関数を定義
// 	modify := func(a []string, f func([]string) []string) []string {
// 		return f(a)
// 	}

// 	// stringのスライスを作成
// 	m := []string{"1st", "2nd" ,"3rd"}

// 	// modify関数の呼び出し
// 	m1 := modify(m, func([]string) []string {
// 		return append(m, m...)
// 	})

// 	fmt.Println(m1)
// }

func main() {
	n := 123 // 変数
	p := &n  // その値のアドレスを取得
	fmt.Println("number", n)
	fmt.Println("pointer", &n)
	fmt.Println("value", *p)
	
}