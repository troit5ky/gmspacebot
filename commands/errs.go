package commands

import (
	"log"

	"github.com/lus/dgc"
)

func connectionTimeOut(ctx *dgc.Ctx, err error) bool {
	log.Println(err)
	embed.Color = 16730698
	embed.Description = `Сервер недоступен ¯\_(ツ)_/¯`
	ctx.RespondEmbed(embed)
	return false
}

func notEnoughPlayers(ctx *dgc.Ctx) bool {
	embed.Description = `Никого нету ¯\_(ツ)_/¯`
	embed.Color = 16730698
	ctx.RespondEmbed(embed)
	return false
}

func notEnoughArgs(ctx *dgc.Ctx) bool {
	embed.Description = "Введи аргументы!"
	ctx.RespondEmbed(embed)
	return false
}
