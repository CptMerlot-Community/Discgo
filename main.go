package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var TOKEN = ""

func init() {
	CreateDiscordSession()
	BuildCommands()
	CreateSlashHandler()
}

func main() {

	Discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := Discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shuting down")

	err = Discord.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

}
