package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	configVar *configStruct = &configStruct{}
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func Config() error {

	file, err1 := ioutil.ReadFile("./config.json")
	if err1 != nil {
		fmt.Println("Error During File Read", err1.Error())
		return err1
	}

	err2 := json.Unmarshal(file, configVar)
	if err2 != nil {
		fmt.Println("Error During Unmarshall Read", err2.Error())
		return err2
	}

	Token = configVar.Token
	BotPrefix = configVar.BotPrefix
	return nil
}
