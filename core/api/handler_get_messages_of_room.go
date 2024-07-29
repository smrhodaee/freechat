package api

import (
	"app/constants"
	"app/models"
	"app/permissions"
	"app/router"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type RoomRequest struct {
	RoomName string       `form:"room_name" json:"room_name"`
	Room     *models.Room `form:"-" json:"-"`
}

func (r *RoomRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	room, err := s.Room.GetByName(r.RoomName)
	if err != nil {
		return err
	}
	if room == nil {
		return router.NotFoundError()
	}
	if !permissions.CanReadFromRoom(router.GetReqInfo(c).User, room) {
		return router.PermissionError()
	}
	r.Room = room
	return nil
}

type GetMessagesOfRoomHandler struct {
	s *store.Store
}

func (h *GetMessagesOfRoomHandler) V1(c *fiber.Ctx) error {
	var req RoomRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	chats, err := h.s.Message.GetMessagesOfRoom(req.Room)
	if err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(
		constants.GetChatsSuccessfully,
		router.Collection(ChatResponse{}, chats),
	))
}
