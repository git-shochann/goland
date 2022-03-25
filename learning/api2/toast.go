package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/ini.v1"
)

var TokenEndpoint = "https://api-identity.infrastructure.cloud.toast.com/v2.0/tokens"

// https://docs.toast.com/ja/Compute/Compute/ja/identity-api/

type RequestBody struct {
	Auth Data `json:"auth"`
}
type Data struct {
	Tenantid            string              `json:"tenantId"` // タグに空白入れるとエラー発生する
	Passwordcredentials Passwordcredentials `json:"passwordCredentials"`
}

type Passwordcredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Body RequestBody // 全スコープ対象に型を宣言 = グローバル変数 // 実際にインスタンス化したBodyをmain関数で使用するため。

func init() {
	cfg, err := ini.Load("../../config.ini")
	if err != nil {
		fmt.Printf("fail to load config file: %v", err)
		os.Exit(1)
	}

	Body = RequestBody{
		Auth: Data{
			Tenantid: cfg.Section("toast").Key("tenantid").String(),
			Passwordcredentials: Passwordcredentials{
				Username: cfg.Section("toast").Key("username").String(),
				Password: cfg.Section("toast").Key("password").String(),
			},
		},
	}
}

func main() {
	fmt.Println(Body.Auth.Passwordcredentials.Username) // sunitaro216@gmail.com
	json, err := json.Marshal(Body)                // JSONに変換
	if err != nil {
		fmt.Printf("fail to encode json: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(json))

	req, err := http.NewRequest("POST", TokenEndpoint, bytes.NewBuffer(json)) // リクエストの作成
	if err != nil {
		fmt.Printf("err")
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("err")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))

}