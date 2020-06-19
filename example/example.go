package main

import (
	send "github.com/salmaptr/telegram-notif"
)

var (
	telegramBotURL = "https://api.telegram.org/bot<token>/sendMessage"
)

func main() {
	id := "<user id>"
	msg := "Hello World"

	// 1. telegramBotURL
	// 2. msg
	// 3. UserID
	send.NotifyTelegram(telegramBotURL, msg, id)
}
