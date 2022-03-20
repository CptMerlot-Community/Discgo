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
