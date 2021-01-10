package commands

import (
	"log"
	"time"

	"github.com/lus/dgc"
)

func Init(router *dgc.Router) {
	log.Printf("Initialize commands...")

	defer log.Printf("Commands initialized!")

	router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Вернёт pong",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 3*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("Нельзя использовать бота так часто!")
		}),
		Handler: Ping,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "echo",
		Description: "Вернёт аргументы",
		Usage:       "echo [a]...",
		Example:     "echo Hello World!",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 3*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("Нельзя использовать бота так часто!")
		}),
		Handler: Echo,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "status",
		Description: "Вернёт статус Garry's Mod сервера",
		Usage:       "status",
		Example:     "status",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 3*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("Нельзя использовать бота так часто!")
		}),
		Handler: Status,
	})
}
