package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/CptMerlot-Community/Discgo/gh"
	"github.com/CptMerlot-Community/Discgo/session"
	"github.com/bwmarrin/discordgo"
)

var gh_client = gh.CreateClient(nil)

func CreateSlashHandler(s *session.Session) {

	if s == nil {
		log.Panicln("Discord Session not setup")
	}
	as := s.GetActiveSession()
	if as != nil {
		as.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if i != nil {
					log.Printf("Handling request %s, Command Name %s", i.ID, i.ApplicationCommandData().Name)
				}
				switch i.ApplicationCommandData().Name {
				case "hello":
					handlerHelloSlash(s, i)
				case "test":
					handlerTestSlash(s, i)
				}
			case discordgo.InteractionMessageComponent:
				log.Printf("%+v", i.MessageComponentData())
				log.Printf("%+v", *i)
				log.Printf("Content - %s", i.Message.Content)
				// i.Message
				switch i.Message.Content {
				case "Yes or No":
					print("correct")
				}
			case discordgo.InteractionModalSubmit:
				switch {
				case strings.HasPrefix(i.ModalSubmitData().CustomID, "sign_up"):
					signUpModalSubmit(s, i)
				case strings.HasPrefix(i.ModalSubmitData().CustomID, "other"):
					fmt.Println("OTHER")
				}

			}
		})
	}

}

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

// Look at this later
//				discordgo.ActionsRow{
//					Components: []discordgo.MessageComponent{
//						discordgo.SelectMenu{
//							CustomID:    "select",
//							Placeholder: "Choose your favorite programming language 👇",
//							Options: []discordgo.SelectMenuOption{
//								{
//									Label:       "test_label",
//									Value:       "test_label",
//									Description: "This is a test",
//									Default:     true,
//								},
//							},
//						},
//					},
//				},
