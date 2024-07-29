package api

import (
	"app/constants"
	"app/router"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type GetRoomsOfUserHandler struct {
	s *store.Store
}

func (h *GetRoomsOfUserHandler) V1(c *fiber.Ctx) error {
	rooms, err := h.s.Room.GetRoomsOfUser(router.GetReqInfo(c).User)
	if err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(
		constants.GetRoomsSuccessfull,
		router.Collection(RoomResponse{}, rooms),
	))
}
