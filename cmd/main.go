package main

import (
	"humo_bot/bot"
	"humo_bot/config"
	"humo_bot/db"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	bot.StartBot()

}
