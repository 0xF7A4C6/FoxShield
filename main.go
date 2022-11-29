package main

import (
	"bot/lib/core/database"
	"bot/lib/core/discord"
	"bot/lib/core/rest"
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

	Err := discord.InitBot("MTA0NjM2NTc2MDc1NjM4Mzc2NA.GF6dGD.r7PleXxOGbM-_FiWbT7aP1d05NHAhheTArFnDg")
	if Err != nil {
		panic(Err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
