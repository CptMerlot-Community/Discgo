package handlers

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func signUpModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {

	data := i.ModalSubmitData()

	v1 := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	_ = data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	fmt.Printf("Value1 - %s\n", v1)
	if !gh_client.CheckValidUser(v1) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please enter a valid Github User",
				Flags:   1 << 6,
			},
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Thank you for registering",
			Flags:   1 << 6,
		},
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(data.CustomID)

	//if !strings.HasPrefix(data.CustomID, "modals_survey") {
	//	return
	//}

	//userid := strings.Split(data.CustomID, "_")[2]
	//_, err = s.ChannelMessageSend(*ResultsChannel, fmt.Sprintf(
	//	"Feedback received. From <@%s>\n\n**Opinion**:\n%s\n\n**Suggestions**:\n%s",
	//	userid,
	//	data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
	//	data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value,
	// ))
	// if err != nil {
	// 	panic(err)
	// }
}
