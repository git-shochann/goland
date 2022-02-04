package mylib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key" : "User1Secret",
	"User2Key" : "User2Secret",
}

// サーバーサイド側 -> データが正しいかどうか確認する
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

// クライアント側 -> リクエストを投げる
func Client() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret)) // sha256(ハッシュ関数)にて, apiSecretをハッシュ化させる
	fmt.Println(h)

	h.Write(data) // データを書き込む
	sign := hex.EncodeToString(h.Sum(nil)) // 
	fmt.Println(sign) // 076b55e7f7e12624b4569f162302f1e36c11fb3a9134889267748de14a84b996

	Server(apiKey, sign, data) // サーバーサイドに投げる
}