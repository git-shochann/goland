package main

import "fmt"

type dog4 struct {
	name string
	group string
	height int
}


// dog4のインスタンス.メソッド名でこの後呼び出す

func(d dog4) describe() string {
	dscr := "私は" + d.group
	dscr += "、名は" + d.name
	return dscr // "私はマルチーズ、名はマル"
}

// dog4のインスタンスを引数に取る
func (d dog4) bigger(d2 dog4) string {
	dscr := d.describe() // {"マル", "マルチーズ", 22}
	fmt.Println("---")
	fmt.Println(dscr) // "私はマルチーズ、名はマル"
	fmt.Println("---")
	dscr += "、私は" + d2.name // {"シバ", "柴犬", 40}

	diff := d.height - d2.height // 22 - 40
	switch {
	case diff > 5:
		dscr += "より大きい"
	case diff < -5:
		dscr += "より小さい"
	default: 
		dscr += "と同じくらいである"
	}
	return dscr
}

func main(){
	maru := dog4{"マル", "マルチーズ", 22}
	shiba := dog4{"シバ", "柴犬", 40}

	// インスタンス.メソッド()
	fmt.Println(maru.bigger(shiba))
}