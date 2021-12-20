package commands

import (
	"gmspacebot/lib/config"
	"strconv"

	"github.com/lus/dgc"
	"github.com/rumblefrog/go-a2s"
)

func Players(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Информация о игроках"

	client, err := a2s.NewClient(config.Option.IP)

	if err != nil {
		connectionTimeOut(ctx, err)
		return
	}

	defer client.Close()

	info, err := client.QueryPlayer()
	if err != nil {
		connectionTimeOut(ctx, err)
		return
	}

	if (int(info.Count)) == 0 {
		notEnoughPlayers(ctx)
		return
	}

	var result string

	for i, player := range info.Players {
		plytimeraw := int(player.Duration)
		s := strconv.Itoa(plytimeraw % 60)
		m := strconv.Itoa(plytimeraw / 60 % 60)
		h := strconv.Itoa(plytimeraw / 3600 % 24)

		if player.Name != "" {
			result = result + strconv.Itoa(i+1) + ". [" + h + ":" + m + ":" + s + "] " + player.Name + "\n"
		} else {
			result = result + strconv.Itoa(i+1) + ". [" + h + ":" + m + ":" + s + "] ..." + "\n"
		}
	}

	embed.Description = result
	ctx.RespondEmbed(embed)
}
