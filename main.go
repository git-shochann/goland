package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// 構造体にメソッドを追加する
// DBに保存するように, txtファイルに保存するメソッド
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// DBから探すようにタイトルを条件に, txtファイルの中身を読み取る
func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplete(w http.ResponseWriter, tmpl string, p *Page) {
}

// インターフェースを引数に取る?
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// view/test
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplete(w, "view", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	// Webサーバーの立ち上げ -> エラー起きたらそのままエラー
	log.Fatalln(http.ListenAndServe(":8080", nil))
}