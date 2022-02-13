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

const APP_ID = ""
const GUILD_ID = ""

var s *session.Session

func init() {
	s = session.CreateDiscordSession(APP_ID, GUILD_ID)
	BuildCommands(s)
	handlers.CreateSlashHandler(s)
}

func main() {
	as := s.GetActiveSession()
	if as == nil {
		log.Panicln("No active discord session")
	}
	as.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	err := as.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shuting down")

	err = as.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

}
