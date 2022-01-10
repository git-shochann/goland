package mylib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Http() {
	resp, _ := http.Get("https://www.supremenewyork.com/shop/all")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	base, _ := url.Parse("https://www.supremenewyork.com/shop/all")
	fmt.Println(base)
}