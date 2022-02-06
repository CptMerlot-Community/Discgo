package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/CptMerlot-Community/Discgo/handlers"
	"github.com/CptMerlot-Community/Discgo/session"

	"github.com/bwmarrin/discordgo"
)

var TOKEN = ""

func init() {
	session.CreateDiscordSession()
	BuildCommands()
	handlers.CreateSlashHandler()
}

func main() {

	session.Discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := session.Discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shuting down")

	err = session.Discord.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

}
