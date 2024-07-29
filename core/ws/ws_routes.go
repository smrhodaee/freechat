package ws

import (
	"app/router"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (w *WS) RegisterRoutes(r *fiber.App) *WS {
	r.Get("/chat",
		router.LoginRequiredMiddleware(),
		websocket.New(w.chatHandler),
	)
	return w
}
