package mylib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Http() {
	// 超簡単なGET
	// resp, _ := http.Get("https://www.supremenewyork.com/shop/all")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	// URLが正しいか認識する
	base, _ := url.Parse("https://www.supremenewyork.com/shop/all")
	fmt.Println(base)

	// URLを生成する
	reference, _ := url.Parse("test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	// GETリクエストを作成
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `testtest`)
	q := req.URL.Query() // クエリストリング
	fmt.Println(q)       // map[a:[1] b:[2]]
	q.Add("c", "3")
	req.URL.RawQuery = q.Encode() // なんだこれは？

	// 実際にアクセスする
	var Client *http.Client = &http.Client{} // なんだこれは？
	resp, _ := Client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// POSTを作成したかったら...
	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
}