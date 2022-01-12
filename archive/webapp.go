// DB使わずにTxtファイルに保存するApp

package archive

import (
	"fmt"
	"io/ioutil"
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

func main() {
	// 構造体の初期化
	// ようやくここで値をセットしてメソッドを呼び出す
	p1 := &Page{Title: "web_application", Body: []byte("Sample")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}