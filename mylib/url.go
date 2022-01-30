package mylib

import (
	"fmt"
	"net/url"
)


func init() {
	
	url := &url.URL{} // ポインタ型で指定 -> 構造体の初期化
	url.Scheme = "https"
	url.Host = "shochann.com"
	fmt.Println(url) // https://shochann.com

	q := url.Query() // type Values map[string][]string を戻り値として返す
	q.Set("q", "sho") // 構造体Valuesが持っているSetメソッド
	fmt.Println(q) // map[q:[sho]]
	a := q.Encode()
	fmt.Println(a) // q=sho
	url.RawQuery = a
	fmt.Println(url) // https://shochann.com?q=sho

	q.Set("q", "たろ")
	a = q.Encode()
	fmt.Println(a) // q=%E3%81%9F%E3%82%8D エンコードすることで日本語の値にも対応させることが出来る

}