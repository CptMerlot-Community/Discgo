package commands

import (
	"github.com/CptMerlot-Community/Discgo/session"
	"github.com/bwmarrin/discordgo"
)

func BuildCommands(s *session.Session) {
	var commands = []*discordgo.ApplicationCommand{}
	if s != nil {
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
			{
				Name:        "test",
				Description: "basic test command",
			},
		}
	}

	if len(commands) > 0 {
		s.RegisterCommands(commands)
	}
}
