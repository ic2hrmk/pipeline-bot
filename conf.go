package main

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Configuration struct {
	//	To connect to Telegram API
	AuthToken *string	`json:"auth_token"`
	//	To connect to chat
	Password *string `json:"password"`
}

const splashScreen = `______ _            _ _             ______       _
| ___ (_)          | (_)            | ___ \     | |
| |_/ /_ _ __   ___| |_ _ __   ___  | |_/ / ___ | |_
|  __/| | '_ \ / _ \ | | '_ \ / _ \ | ___ \/ _ \| __|
| |   | | |_) |  __/ | | | | |  __/ | |_/ / (_) | |_
\_|   |_| .__/ \___|_|_|_| |_|\___| \____/ \___/ \__|
		| |
		|_|                                          `

const description  = "=====  Telegram bot for pipeline notification =====\n"

func NewConfiguration(fileName string) (conf *Configuration, err error) {
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonData, &conf)

	if err != nil {
		return
	}

	if conf.Password == nil {
		conf.Password = new(string)
	}

	return
}

func (conf *Configuration) IsValid() (valid bool, err error) {
	if conf == nil {
		err = errors.New("empty configuration")
		return
	}

	if conf.AuthToken == nil {
		err = errors.New("auth token required")
		return
	}

	if conf.AuthToken == nil {
		err = errors.New("auth token required")
		return
	}

	valid = true

	return
}
