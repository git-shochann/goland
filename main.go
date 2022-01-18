package main

import (
	"html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html") // 外部ファイルを取り込む
	t.Execute(w, p) // そのテンプレートに対して*Page情報を渡して実行
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // view/test -> 変数.構造体.フィールド名で取得
	p, err := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	if err != nil {
		http.Redirect(w, r, "/edit"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} // エラーの場合だったとしても, 新しいページを作ってrenderTemplateを呼び出す
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body  := r.FormValue("body") // saveの際にやってくる実際のフォームを取得
	p     := &Page{Title: title, Body: []byte(body)}
	err   := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500を返す
		return // 処理を抜ける
	}
	http.Redirect(w, r, "/view/", http.StatusFound)
}
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil)) 	// Webサーバーの立ち上げ -> エラー起きたらそのままエラー出力
}