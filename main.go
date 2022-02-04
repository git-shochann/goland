// 常に実行できるように

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {

	mac := hmac.New(sha256.New, []byte("hello world"))
	fmt.Println()
}