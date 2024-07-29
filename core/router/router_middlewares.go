package router

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

func authMiddleware(store *store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)
		if token == "" {
			token = c.Query("token")
		}
		if len(token) == configs.TokenLength {
			user, err := store.Token.GetUser(token)
			if err != nil {
				return err
			}
			if user == nil {
				return fiber.NewError(fiber.StatusUnauthorized, constants.AuthorizationFailed)
			}
			c.Locals(configs.LocalReqInfo, models.NewReqInfo(token, user))
		}
		return c.Next()
	}
}

func LoginRequiredMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqInfo := GetReqInfo(c)
		if reqInfo == nil {
			return fiber.NewError(fiber.StatusUnauthorized, constants.LoginRequired)
		}
		if !reqInfo.IsActive {
			return fiber.NewError(fiber.StatusUnauthorized, constants.UserNotActive)
		}
		return c.Next()
	}
}