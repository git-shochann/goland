package mylib

import "fmt"

type Mydata struct {
	// 変数 + 型
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("Name:", md.Name)
	fmt.Println("Data", md.Data)	
}

func init() {
	// 構造体の初期化
	taro := Mydata{
		"Hanako", []int{10,20,30,40,50},
	}
	// taroという変数は構造体Mydataから初期化している = 関数を所有しているので以下のように呼び出すことが可能
	taro.PrintData()
}

// 構造体型の中に構造体型

type User struct {
	Name string
	Email string
	Profile Profile // Profileフィールドを以下のProfile型として設定
}

type Profile struct {
	Age int
}

func Do() {

	user := User{ 
		Name: "Taro", 
		Email: "t0825.st@gmail.com",
		Profile: Profile{Age: 24},
  }
  fmt.Printf("%v, %T\n", user, user) // %v -> value, %T -> Type
  // {Taro t0825.st@gmail.com {24}}, main.User 型はmainのUser型

}