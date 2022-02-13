package session

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var TOKEN_PATH = os.Getenv("TOKEN_PATH")

type Session struct {
	AppID          string
	GuildID        string
	discordSession *discordgo.Session
}

func (s Session) GetActiveSession() *discordgo.Session {
	return s.discordSession
}

func (s Session) RegisterCommands(commands []*discordgo.ApplicationCommand) {

	for _, c := range commands {
		_, err := s.discordSession.ApplicationCommandCreate(s.AppID, s.GuildID, c)
		if err != nil {
			log.Printf("Command %s had error %s", c.Name, err.Error())
		}
	}

}

func CreateDiscordSession(appID, guildID string) *Session {
	s := &Session{
		AppID:   appID,
		GuildID: guildID,
	}

	f, err := os.ReadFile(TOKEN_PATH)
	if err != nil {
		log.Panicln(err)
	}

	TOKEN := string(f)

	s.discordSession, err = discordgo.New(fmt.Sprintf("Bot %s", TOKEN))
	if err != nil {
		log.Panicln(err)
	}

	return s

}
