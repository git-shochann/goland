package main

import "fmt"

type someapp struct {
	username string
	use int
	isopen bool
}

func (app *someapp) open() {
	if app.isopen {
		fmt.Printf("%sさんのアプリは既に開いています。\n", app.username)
	} else {
		app.use++
		if app.use > 2 {
			fmt.Printf("%sさんの試用期間は終了です。ご購入をご検討下さい。\n", app.username)
		} else {
			fmt.Printf("こんにちは%sさん。%d回目の使用になります。\n", app.username, app.use)
			app.isopen = true
		}
	}
}

func (app *someapp) close() {
	if app.isopen {
		fmt.Printf("さようなら%sさん。アプリを終了します。\n", app.username)
		app.isopen = false
	}
}

// 構造体のインスタンスを戻り値で返す
func newapp(username string) someapp {
	fmt.Printf("ようこそ%sさん。初めてのご利用になります。\n", username)
	return someapp{username, 1, true} // ここで構造体をインスタンス化	
}

func main() {
	fmt.Println("りんごさんがアプリを利用開始しました。")
	rapp := newapp("りんご")
	rapp.close()

	fmt.Println(rapp) // {りんご 1 false}

	fmt.Println("りんごさんが再度Appを開きました。")
	rapp.open()
	rapp.close()
	fmt.Println(rapp) // {りんご, 2, false}

	fmt.Println("みかんさんがアプリを使用開始しました。")
	mapp := newapp("みかん")
	mapp.close()

	fmt.Println("みかんさんが再度Appを開きました。")
	mapp.open()

	fmt.Println("りんごさんが3回目を開きます。")
	rapp.open()

	fmt.Println("みかんさんが3回目を開きます。")
	mapp.open()
	
}