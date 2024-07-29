package api

import (
	"app/constants"
	"app/router"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type LogoutHandler struct {
	s *store.Store
}

func (h *LogoutHandler) V1(c *fiber.Ctx) error {
	if err := h.s.Token.Delete(router.GetReqInfo(c).Token); err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(constants.LogoutSuccessfull))
}
