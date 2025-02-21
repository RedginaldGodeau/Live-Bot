package discord

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	*discordgo.Session
}

func NewBot(token string) (*DiscordBot, error) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &DiscordBot{bot}, nil
}

func (bot *DiscordBot) Start() error {
	err := bot.Open()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println("Bot Discord en ligne.")
	select {}
}

func (bot *DiscordBot) AddCommand(command string, handler func(s *discordgo.Session, m *discordgo.MessageCreate)) {
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		content := m.Content
		if content == "" {
			chanMsg, err := s.ChannelMessages(m.ChannelID, 1, "", "", m.ID)
			if err != nil {
				log.Println("unable to get messages:", err)
				return
			}
			content = chanMsg[0].Content
		}

		if strings.HasPrefix(content, command) {
			handler(s, m)
		}
	})
}
