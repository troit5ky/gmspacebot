package commands

import (
	"github.com/lus/dgc"
)

func Ping(ctx *dgc.Ctx) {
	ctx.RespondText(":ping_pong: **Pong!**")
}
