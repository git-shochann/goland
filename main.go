// 常に実行できるように

package main

import "fmt"

type User struct {
	Name string
	Email string
	Profile Profile // Profileフィールドを以下のProfile型として設定
}

type Profile struct {
	Age int
}

func main() {

	user := User{ 
		Name: "Taro", 
		Email: "t0825.st@gmail.com",
		Profile: Profile{Age: 24},
  }
  fmt.Printf("%v, %T\n", user, user) // %v -> value, %T -> Type
  // {Taro t0825.st@gmail.com {24}}, main.User 型はmainのUser型

}