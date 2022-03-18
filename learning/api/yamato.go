package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var endpoint = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"


func main() {
	query := url.Values{} // 自分でフィールドを設定 type Values map[string][]string == データ型は、キーがstring, バリューがstringのスライス
	query.Add("number00", "1")
	query.Add("number01", "4556-7720-1976")
	fmt.Println(&query) // なぜアドレスでないとダメなのか？
	body := strings.NewReader(query.Encode()) // io.Readerを満たすインターフェースの用意？
	req, err := http.NewRequest("POST", endpoint, body) // リクエスト作成するだけ
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(req)
	fmt.Println(req.Header) // map[]
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // http.Request構造体のHeaderフィールドを置き換え
	fmt.Println(req.Header) // map[Content-Type:[application/x-www-form-urlencoded]]

	// http.Doメソッドはhttp.Client構造体の持っているメソッドであるので構造体の初期化が必要？ ここはポインタ型にしたほうがいいの？
	client := http.Client{}
	res, err := client.Do(req)
	fmt.Println(res.StatusCode)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close() // http.Response構造体 -> フィールドBody io.ReadCloserインターフェース -> Reader Closerインターフェース -> Closeメソッドを実装している ここでのCloseメソッドを使おうと思える流れは？ そもそもこんな感じでやっているのか？

	// この後のBody読み取りはどうする？ ググってioutil.ReadAll()で読み取りが必要だと確認した。(func ReadAll(r io.Reader) ([]byte, error) io.Readerインターフェースだからhttp.Doメソッドの戻り値の構造体であるhttp.Responseの構造体のフィールドであるBody io.ReadCloser に対して使えるのか？ そもそもここまで考えているか？

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(data))

	// 今回はスクレイピング目的なので上記は不要

	doc, err := goquery.NewDocumentFromReader(res.Body) // Document構造体の色々な情報を使ったり、メソッドを使うためにパースする パースするとはこういった意味で捉えてOKか
	if err != nil {
		log.Fatalln(err)
	}
	// result := doc.Find(".data.arrow.sp-only").Each(func(i int, s *goquery.Selection){ // 関数の引数に関数？ ここではどういうふうな流れで使ってる？
	// 	fmt.Println(s.Text()) // selection構造体のText()メソッド？ // ▶ が返ってくる -> 
	// })
    
	elem := doc.Find(".js-tracking-detail")
	fmt.Println(elem.Text())

}

// ①
// httpのNewRequestを使いたいときにドキュメントを見て、 
// 引数にPOST, endpoint, io.ReaderとあったときにReaderインターフェースを満たしている構造体が必要なんだなあ考えて、
// その後Readerインターフェースを満たしてる構造体ってどう探せばいいのでしょうか？ ここでの頭の動かし方はどんな感じでしょうか？


// ②
// 基本的なHTML, CSS？。Textメソッドは、子孫を含む、一致した要素のセット内の各要素の結合されたテキストコンテンツを取得するとは？
// セレクタ？

// ③
// func (s *Selection) Find(selector string) *Selection
// この引数に取るselectorとはどう指定するのか？ ドキュメントを見ても書いていない。

// とりあえず一件ステータスをチェックして、それを取得できれば一旦OK!