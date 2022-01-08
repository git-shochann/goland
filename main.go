package main

import (
	"goland/mylib"
	"strconv"
)

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
c// 	} else {
// 		return // main関数を終了するという意味
// 	}
// 	switch {
// 		case n%2 == 0:
// 			fmt.Println("偶数です")
// 		case n%2 == 1:
// 			fmt.Println("奇数です")
// 	}
// }

func main() {
	x := mylib.Input("Type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Println("1から" + x + "の偶数の合計は")
	} else {
		return
	}
	t := 0
	c := 0
	for {
		c++
		if c%2 == 1 {
			continue
		}
		if c > n {
			break
		}
		t += c
	}
	fmt.Println(t, "です")
}