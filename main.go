package main

import (
	"bot/lib/core/database"
	"bot/lib/core/discord"
	"bot/lib/core/rest"
	"bot/lib/utils"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	database.InitDatabase()
	go rest.HandleRequests()

	Err := discord.InitBot(utils.TestToken)
	if Err != nil {
		panic(Err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
