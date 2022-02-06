package handlers

import (
	"fmt"
	"log"

	"github.com/CptMerlot-Community/Discgo/session"
	"github.com/bwmarrin/discordgo"
)

func CreateSlashHandler() {
	if session.Discord == nil {
		log.Panicln("Discord Session not setup")
	}
	session.Discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.ApplicationCommandData().Name {
		case "hello":
			handlerHelloSlash(s, i)
		}
	})
}

func handlerHelloSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "echo":
		handleEchoSubHello(s, i)
	case "user":
		handleUserSubHello(s, i)
	case "hyb":
		handleHYBSubHello(s, i)
	}
}

func handleEchoSubHello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := i.ApplicationCommandData().Options[0].Options[0].Value
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Discgo Says - %s", content),
		}})
	if err != nil {
		fmt.Println(err)
	}
}
func handleUserSubHello(s *discordgo.Session, i *discordgo.InteractionCreate) {

	user := i.User

	if user == nil {
		user = i.Member.User
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Grettings %s", user.Username),
		}})
	if err != nil {
		fmt.Println(err)
	}
}
func handleHYBSubHello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "HELL YEA BROTHER!!111one",
		}})
	if err != nil {
		fmt.Println(err)
	}
}
