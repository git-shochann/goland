package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// 同じtemplateであっても、処理毎に毎回templateの読み込みが行われているのでキャッシュする
// 戻り値で構造体*Templateを返す -> 変数templetesはその構造体が持っているメソッドを使用することが出来る
var templetes =  template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, _ := template.ParseFiles(tmpl + ".html") // 外部ファイルを取り込む
	// t.Execute(w, p) // そのテンプレートに対して*Page情報を渡して実行
	err := templetes.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):] // view/test -> 変数.構造体.フィールド名で取得
	p, err := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	if err != nil {
		http.Redirect(w, r, "/edit"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} // エラーの場合だったとしても, 新しいページを作ってrenderTemplateを呼び出す
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	body  := r.FormValue("body") // saveの際にやってくる実際のフォームを取得
	p     := &Page{Title: title, Body: []byte(body)}
	err   := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500を返す
		return // 処理を抜ける
	}
	http.Redirect(w, r, "/view/", http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$") // 日本語を読むように書く

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r* http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path) // validPathと一致するか確認
		if m == nil {
			http.NotFound(w,r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatalln(http.ListenAndServe(":8080", nil)) 	// Webサーバーの立ち上げ -> エラー起きたらそのままエラー出力
}