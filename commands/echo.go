package commands

import (
	"github.com/lus/dgc"
)

func Echo(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Echo"

	if ctx.Arguments.Amount() < 1 {
		notEnoughArgs(ctx)
	} else {
		embed.Description = ctx.Arguments.Raw()
		ctx.RespondEmbed(embed)
	}
}
