package main

import (
	"pipeline-bot/server"
	"pipeline-bot/telegram"

	"fmt"
	"log"
	"flag"
	"net/http"
)

var port string
var fileName string

func init() {
	flag.StringVar(&port, "port", "8080", "port to send web-hooks")
	flag.StringVar(&fileName, "conf", "conf.json", "configuration file")
	flag.Parse()
}

func startUpMessage() {
	fmt.Println(splashScreen)
	fmt.Println(description)

	fmt.Println("WebHook handler started at port: " + port)
	fmt.Println("Configuration file: " + fileName + "\n")
}

func initTelegramBot(conf *Configuration) (err error) {
	return telegram.NewBot(*conf.AuthToken)
}

func initHttpReceiver(conf *Configuration) (err error) {
	http.HandleFunc("/hooks", server.WebHookHandler)
	return http.ListenAndServe(":" + port, nil)
}

func main() {
	startUpMessage()

	conf, err := NewConfiguration(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	if ok, err := conf.IsValid(); !ok {
		log.Fatal(err.Error())
	}

	//	Initialize modules
	err = initTelegramBot(conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = initHttpReceiver(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}