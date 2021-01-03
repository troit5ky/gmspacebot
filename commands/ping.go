package commands

import (
	"strconv"
	"time"

	"github.com/lus/dgc"
)

func Ping(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Pong!"

	Parse, _ := ctx.Event.Timestamp.Parse()
	Timestamp := Parse.Unix()
	Timenow := time.Now().Unix()
	Result := strconv.FormatInt(Timenow-Timestamp+3, 10)

	embed.Description = "Time: `" + Result + "ms`"
	ctx.RespondEmbed(embed)
}
