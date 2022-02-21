package main

// import "fmt"

// type dog3 struct {
// 	name string
// 	group string
// 	height int
// }

// type bird struct {
// 	name string
// 	group string
// 	length int
// }

// func (d dog3) describe() string {
// 	dscr := "私は" + d.group
// 	dscr += "、名は" + d.name
// 	return dscr // "私はポメラニアン、名はポメ"
// }

// func (b bird) describe() string {
// 	dscr := "私は" + b.group
// 	dscr += "、名は" + b.name
// 	return dscr
// }

// func main() {
// 	pome := dog3{"ポメ", "ポメラニアン", 25} // 構造に従って、値を与える -> インスタンスの作成
// 	meji := bird{"メジ", "スズメ", 12}
// 	fmt.Println(pome.describe()) // そのまま戻り値を出力
// 	fmt.Println(meji.describe())
// }