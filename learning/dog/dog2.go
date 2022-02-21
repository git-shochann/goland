package main

// import "fmt"

// type dog2 struct {
// 	name string
// 	group string
// 	height int
// }

// func biggest(dogs []dog2) dog2 {

// 	biggest := dogs[0] // {ポメ ポメラニアン 25} 基準を設定する

// 	for i:=0; i < len(dogs); i ++ {
// 		fmt.Println(dogs[i].height) // 25
// 		fmt.Println(biggest.height) // 25
// 		if dogs[i].height > biggest.height { // 25 > 25 -> 22 > 25 -> 40 > 25
// 			biggest = dogs[i]
// 		}
// 	}
// 	return biggest // ここでbiggestに代入されているのを戻り値として返す
// }

// func main() {
// 	pome := dog2{"ポメ", "ポメラニアン", 25}
// 	maru := dog2{"マル", "マルチーズ", 22}
// 	shiba := dog2{"シバ", "柴犬", 40}

// 	dogs := []dog2{pome, maru, shiba} // dogインスタンスのスライス

// 	fmt.Printf("%sさんが最大です\n", biggest(dogs).name)
// }