package main

import "fmt"

func main() {

	price := []int{10, 20, 30, 40, 50}
	items := []string{"apple", "banana", "cherry", "orange", "kiwi"}

	sum := 0

	for i := 0; i < len(price); i++ {
		sum += price[i]
		fmt.Printf("%s is ¥%d , 小計 ¥%d\n", items[i], price[i], sum)
	}

	fmt.Printf("合計金額は¥%dです。", sum)

}
