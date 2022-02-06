package session

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

const APP_ID = ""
const GUILD_ID = ""

var Discord *discordgo.Session

func CreateDiscordSession() {

	f, err := os.ReadFile("C:\\Users\\cptme\\SUPER_SECRET_TOKEN_BOI")
	if err != nil {
		log.Panicln(err)
	}

	TOKEN := string(f)

	Discord, err = discordgo.New(fmt.Sprintf("Bot %s", TOKEN))
	if err != nil {
		log.Panicln(err)
	}

}
