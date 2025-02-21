package route

import (
	"backend/ent/gen"
	"backend/ent/gen/liveshow"
	"backend/ent/gen/upload"
	"backend/internal/utils"
	"backend/pkg/application"
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func DiscordCommand(app *application.App) {
	bot := app.DiscordBot

	bot.AddCommand("$ping", func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "Hello World")
	})

	bot.AddCommand("$livebot", func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if len(m.Attachments) == 0 {
			s.ChannelMessageSend(m.ChannelID, "Aucun fichier envoyé")
			return
		}

		fileContent, err := utils.UploadDiscordFile(m.Attachments[0].URL)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Une erreur est survenu: "+err.Error())
			return
		}

		contents := strings.Split(m.Content, " ")
		legend := ""
		if len(contents) > 0 {
			legend = strings.Join(contents[1:], " ")
		}

		err = utils.WithTx(context.TODO(), &app.DB.Client, func(tx *gen.Tx) error {
			typeFileDB := upload.FileTypeImage
			if fileContent.FileType == "video" {
				typeFileDB = upload.FileTypeVideo
			}
			upload, err := tx.Upload.Create().SetFilePath(fileContent.FileName).SetName(fileContent.FileName).SetFileType(typeFileDB).Save(context.TODO())
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.LiveShow.Create().SetDuration(fileContent.Duration).SetLegend(legend).SetUploadID(upload.ID).Exec(context.TODO())
			if err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Une erreur est survenu: "+err.Error())
			return
		}

		count, err := app.DB.LiveShow.Query().Where(liveshow.ViewedEQ(false)).Count(context.TODO())
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Une erreur est survenu: "+err.Error())
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Votre contenu à été ajouté à la pile: %d", count))
	})
}
