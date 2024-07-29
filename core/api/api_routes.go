package api

import (
	"app/configs"
	"app/router"
	"app/services/ratelimit"
	"app/store"
	"app/ws"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type API struct {
	ChaptaHandler
	LoginOrRegisterHandler
	LogoutHandler
	CreateGroupHandler
	CreateDirectHandler
	GetMessagesOfRoomHandler
	CreateMessageHandler
	DeleteMessageHandler
	GetImageMessageHandler
	GetRoomsOfUserHandler
	FindUsersHandler
	ChaptaMiddleware
	RatelimitMiddleware
}

func NewAPI(log *logrus.Logger, store *store.Store, ratelimit *ratelimit.Service, ws *ws.WS) *API {
	return &API{
		ChaptaHandler{store, ratelimit},
		LoginOrRegisterHandler{store},
		LogoutHandler{store},
		CreateGroupHandler{store},
		CreateDirectHandler{store},
		GetMessagesOfRoomHandler{store},
		CreateMessageHandler{store, ws},
		DeleteMessageHandler{store, ws},
		GetImageMessageHandler{store},
		GetRoomsOfUserHandler{store},
		FindUsersHandler{store},
		ChaptaMiddleware{store, ratelimit},
		RatelimitMiddleware{ratelimit},
	}
}

func (a *API) Register(app *fiber.App) {
	app.Get("/chapta.png",
		a.RatelimitMiddleware.ByIP(configs.ChaptaRatelimit),
		a.ChaptaHandler.V1,
	)

	app.Post("/auth/login", a.ChaptaMiddleware.Handler, a.LoginOrRegisterHandler.V1)

	app.Use(router.LoginRequiredMiddleware())
	app.Get("/auth/logout", a.LogoutHandler.V1)
	app.Post("/rooms/direct", a.CreateDirectHandler.V1)
	app.Post("/rooms/group", a.CreateGroupHandler.V1)
	app.Post("/rooms/msgs", a.GetMessagesOfRoomHandler.V1)
	app.Post("/msgs", a.CreateMessageHandler.V1)
	app.Post("/delete-message", a.DeleteMessageHandler.V1)
	app.Post("/get-image-chat", a.GetImageMessageHandler.V1)
	app.Get("/users/rooms", a.GetRoomsOfUserHandler.V1)
	app.Post("/users", a.FindUsersHandler.V1)
}
