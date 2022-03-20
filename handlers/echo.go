package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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

func handlerTestSlash(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "sign_up_" + i.Interaction.Member.User.ID,
			Title:    "Sign Up",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:  "gh_user_" + i.Interaction.Member.User.ID,
							Label:     "Github Username",
							Style:     discordgo.TextInputShort,
							Required:  true,
							Value:     "Github Username",
							MaxLength: 100,
							MinLength: 2,
						},
					},
				},
				discordgo.ActionsRow{
					//TODO: Figureout why required false is not making the compoent false
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID: "twitch_user_" + i.Interaction.Member.User.ID,
							Label:    "Twitch Username",
							Style:    discordgo.TextInputShort,
							Value:    "Twitch",
							Required: false,
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}
}
