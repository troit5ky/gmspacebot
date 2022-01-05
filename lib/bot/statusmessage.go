package bot

import (
	"fmt"
	"gmspacebot/lib/config"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
)

var session *discordgo.Session
var playerList string
var embed = &discordgo.MessageEmbed{
	Title: "Подожди минуту...",
}
var embedErr = &discordgo.MessageEmbed{
	Title: "Превышено время ожидания ответа",
	Color: 0,
}

func StatusSendMsg() {
	msg, err := session.ChannelMessageSendEmbed(config.Option.StatusChID, embed)
	if err != nil {
		log.Fatal(err)
	}

	config.Option.StatusChMsgID = msg.ID
	config.Rewrite(*config.Option)
}

func errTimeOut() {
	session.ChannelMessageEditEmbed(config.Option.StatusChID, config.Option.StatusChMsgID, embedErr)
}

func updateStatus() {
	client, err := a2s.NewClient(config.Option.IP)
	if err != nil {
		log.Println(err)
		errTimeOut()
		return
	}
	defer client.Close()

	query, err := client.QueryInfo()
	if err != nil {
		log.Println(err)
		errTimeOut()
		return
	}
	servername := query.Name
	gamemode := query.Game
	plyCount := query.Players
	maxPlys := query.MaxPlayers
	ip := config.Option.IP
	mapnow := query.Map

	desc := fmt.Sprintf("Режим: `%s`\nКарта: `%s`\nИгроков: `%d из %d`\nIP: `%s`", gamemode, mapnow, plyCount, maxPlys, ip)

	// если игроков != 0, то формируем список
	if plyCount != 0 {
		playerList = ""
		queryPly, _ := client.QueryPlayer()

		for k, ply := range queryPly.Players {
			if ply.Name != "" {
				playerList = playerList + fmt.Sprintf("%d. %s\n", k+1, ply.Name)
			} else {
				playerList = playerList + fmt.Sprintf("%d. %s\n", k+1, "...")
			}
		}

		field := &discordgo.MessageEmbedField{
			Name:   "Сейчас на сервере",
			Value:  playerList,
			Inline: false,
		}

		embed.Fields[0] = field
	} else {
		field := &discordgo.MessageEmbedField{
			Name:   "Сейчас на сервере",
			Value:  `Никого нету ¯\_(ツ)_/¯`,
			Inline: false,
		}

		embed.Fields[0] = field
	}
	// конец формирования

	embed.Title = servername
	embed.Description = desc
	embed.Color = 4915018
	embed.Footer = &discordgo.MessageEmbedFooter{Text: "обновлено " + time.Now().Format(time.UnixDate)}
	session.ChannelMessageEditEmbed(config.Option.StatusChID, config.Option.StatusChMsgID, embed)
}

func StatusInit(s *discordgo.Session) {
	session = s

	if config.Option.StatusChMsgID == "null" {
		StatusSendMsg()
	}

	embed.Fields = append(embed.Fields, nil)

	updEvent := time.NewTicker(60 * time.Second)

	for range updEvent.C {
		updateStatus()
	}
}
