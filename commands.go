package main

import (
	"log"

	"github.com/CptMerlot-Community/Discgo/session"
	"github.com/bwmarrin/discordgo"
)

func BuildCommands() {
	var commands = []*discordgo.ApplicationCommand{}
	if session.Discord != nil {
		commands = []*discordgo.ApplicationCommand{
			{
				Name:        "hello",
				Description: "Say hello Discgo",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "echo",
						Description: "Echo back what the user said",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "string",
								Description: "String to echo back to user",
								Required:    true,
							},
						},
					},
					{
						Name:        "user",
						Description: "say hello to the user",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "hyb",
						Description: "HYB BB!!",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
				},
			},
		}
	}

	if len(commands) > 0 {
		registerCommands(commands)
	}
}

func registerCommands(commands []*discordgo.ApplicationCommand) {

	for _, c := range commands {
		_, err := session.Discord.ApplicationCommandCreate(session.APP_ID, session.GUILD_ID, c)
		if err != nil {
			log.Printf("Command %s had error %s", c.Name, err.Error())
		}
	}

}
