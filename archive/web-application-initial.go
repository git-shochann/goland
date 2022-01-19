// DB使わずにTxtファイルに保存するApp

package archive

import (
	"fmt"
	"io/ioutil"
)

type Paper struct {
	Title string
	Body  []byte
}

// 構造体にメソッドを追加する
// DBに保存するように, txtファイルに保存するメソッド
func (p *Page) savingPage() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// DBから探すようにタイトルを条件に, txtファイルの中身を読み取る
func loadingPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func init() {
	// 構造体の初期化
	// ようやくここで値をセットしてメソッドを呼び出す
	p1 := &Page{Title: "web_application", Body: []byte("Sample")}
	p1.savingPage()

	p2, _ := loadingPage(p1.Title)
	fmt.Println(string(p2.Body))
}