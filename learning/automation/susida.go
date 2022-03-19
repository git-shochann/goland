package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	defer driver.Stop()
	if err := driver.Start(); err != nil {
		fmt.Printf("failed to start driver. %s \n", err)
	}
	page, err := driver.NewPage()
	if err != nil {
		fmt.Printf("failed to open a new page. %s \n", err)
	}

	if err := page.Navigate("http://typingx0.net/sushida/play.html"); err != nil {
		fmt.Printf("failed to navigate to susida. %s \n", err)
	}

	fmt.Println("Please select mode to play")

	time.Sleep(10 * time.Second)

}