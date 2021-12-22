package constant

import "fmt"

// constはグローバルに書くことが多い, 他のgoファイルから呼ぶとき大文字でスタートする
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func put() {
	fmt.Println(Pi, Username, Password)
}