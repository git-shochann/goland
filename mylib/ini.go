package mylib

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(), // 8080の取得して構造体に設定
		DbName: cfg.Section("db").Key("name").MustString("example.sql"), // デフォルト設定
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func init() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port) // int 8080
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName) // string stockdata.sql
	fmt.Printf("%T %v\n", Config.SQLDriver,  Config.SQLDriver) // string sqlite3
}