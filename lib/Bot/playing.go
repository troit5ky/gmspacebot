package Bot

import (
	config "../Config"
	"github.com/bwmarrin/discordgo"
)

func playing(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, config.Option.Status)
}
