package main

import (
	"fmt"

	"github.com/santosh/discord-ping/bot"
	"github.com/santosh/discord-ping/config"
)

func main() {
	err := config.Config()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})
}
