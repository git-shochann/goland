package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var endpoint = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"

// func SendToLine() {

// }

func main() {

	for {

		query := url.Values{} // 自分でフィールドを設定 type Values map[string][]string == データ型は、キーがstring, バリューがstringのスライス
		query.Add("number00", "1")
		query.Add("number01", "455677208976")
		fmt.Println(&query)                                 // アドレスと普通の変数どっち？
		body := strings.NewReader(query.Encode())           // io.Readerを満たすインターフェースの用意？
		req, err := http.NewRequest("POST", endpoint, body) // リクエスト作成するだけ
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req)
		fmt.Println(req.Header)                                             // map[]
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // http.Request構造体のHeaderフィールドを置き換え
		fmt.Println(req.Header)                                             // map[Content-Type:[application/x-www-form-urlencoded]]

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

		doc, err := goquery.NewDocumentFromReader(res.Body) // Document構造体の色々な情報を使ったり、メソッドを使うためにパースする パースするとはこういった意味で捉えてOKか
		if err != nil {
			log.Fatalln(err)
		}

		new := doc.Find(".js-tracking-detail").Text()

		fmt.Println(new)

		file, err := os.Open("text.txt")

		var old string
		if err != nil {
			old += ""
		} else {
			data, err := ioutil.ReadAll(file)
			if err != nil {
				os.Exit(1)
			}
			old += string(data)
		}

		defer file.Close()

		if new == old {
			fmt.Println("変化がありません。")
		} else {
			if err := ioutil.WriteFile("text.txt", []byte(new), 0700); err != nil {fmt.Println("error!")}
			fmt.Println("変更がありました。")
			fmt.Printf("伝票番号 %v %v \n", query["number01"], new)
			break
		}

		fmt.Println("5秒後に繰り返します")
		time.Sleep(5 * time.Second)
		// elem.Each(func(i int, s *goquery.Selection) { // 関数の引数に関数？ ここではどういうふうな流れで使ってる？
		// 	fmt.Printf("伝票番号 %v %v\n", query["number01"], s.Text())
		// })
	}

}

// ①
// httpのNewRequestを使いたいときにドキュメントを見て、
// 引数にPOST, endpoint, io.ReaderとあったときにReaderインターフェースを満たしている構造体が必要なんだなあ考えて、
// その後Readerインターフェースを満たしてる構造体ってどう探せばいいのでしょうか？ ここでの頭の動かし方はどんな感じでしょうか？

// ②
// func (s *Selection) Find(selector string) *Selection
// この引数に取るselectorとはどう指定するのか？ ドキュメントを見ても書いていない。

// とりあえず一件ステータスをチェックして、それを取得できれば一旦OK!
