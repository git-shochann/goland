package mylib

import (
	"encoding/json"
	"fmt"
	"log"
)

func init() {

	// JSONデータの生成
	jsonData1 := []byte(`{
		"ID" : 1,
	  	"Name" : "Alex"
    }`)

	type Person struct {
		ID int
		Name string
    }

	var person1 Person

	fmt.Printf("%T, %v\n", person1, person1) // {0, } -> 初期値である 0 と 空文字
	fmt.Printf("%T, %v\n", &person1, &person1) // &{0, } -> ポインタ変数

	// JSONデータのマッピング(第2引数にポインタを渡す)
	err := json.Unmarshal(jsonData1, &person1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("---------------")
	fmt.Printf("%T, %v\n", person1, person1) // {0, } -> 初期値である 0 と 空文字
	fmt.Printf("%T, %v\n", person1, person1) // &{0, } -> ポインタ変数}

	fmt.Printf("%+v\n", person1)

	fmt.Println("---------------")

	// 配列型のJSONの場合
	jsonData2 := []byte(`["Go", "Flutter"]`)
	var array2 []string // マッピングしたい型を決定
	_ = json.Unmarshal(jsonData2, &array2)
	fmt.Printf("%v\n", array2) // [Go Flutter]

	main2()
}

func main2() {

	type People struct {
		ID int
		Name string
	}
	
	people1 := People{ID: 1, Name:"Sho"}
	fmt.Printf("%v", people1) // {ID:1 Name:Alisha}
	
	// JSONに変換する
	returnValue1, _ := json.Marshal(people1)
	
	fmt.Println(returnValue1) // byte型スライスが返る [123 34 73 68 34 58 49 44 34 78 97 1...]
	
	fmt.Printf("%s\n", returnValue1) // string型でキャストしてから出力 -> {"ID":1,"Name":"Alisha"}
}