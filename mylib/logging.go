package mylib

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string){
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 開き方 -> 第二引数どんなモードで開くか
	multiLogFile := io.MultiWriter(os.Stdout, logfile) // 出力先を決定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 日付, 時間, 短いファイル名 (フラグ -> 目印)
	log.SetOutput(multiLogFile) // io.Writerを引数にとって, 実際に書き込む？
}

func init() {
	LoggingSetting("test.log")
	_, err := os.Open("aaa.txt")
	if err != nil {
		log.Fatalln("error!", err)
	}
}