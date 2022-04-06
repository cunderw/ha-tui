package main

import (
	"flag"
	"log"
	"os"

	"github.com/cunderw/ha-tui/internal/ui"
)

func main() {
	file, err := os.OpenFile("ha-tui.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	urlFlag := flag.String("url", "", "Home Assistant URL")
	tokenFlag := flag.String("token", "", "Long lived access token")

	flag.Parse()

	url := *urlFlag
	token := *tokenFlag

	log.Println(url)
	log.Println(token)

	app := ui.StartApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
		panic(err)
	}

}
