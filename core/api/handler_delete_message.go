package api

import (
	"app/models"
	"app/permissions"
	"app/router"
	"app/store"
	"app/ws"

	"github.com/gofiber/fiber/v2"
)

type DeleteMessageHandler struct {
	s  *store.Store
	ws *ws.WS
}

func (h *DeleteMessageHandler) Register(r fiber.Router, path string) {
}

func (h *DeleteMessageHandler) V1(c *fiber.Ctx) error {
	var req DeleteMessageRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	if req.Message.Type == models.MessageTypeImage {
		if err := h.s.File.DeleteByID(req.Message.ValueID); err != nil {
			return err
		}
	}
	if err := h.s.Message.DeleteByID(req.Message.ID); err != nil {
		return err
	}
	h.ws.Broadcast(ws.NewDeleteMessageResponse(req.Message.ID), req.Message.Room.Members.GetUsernames())
	return c.SendStatus(fiber.StatusNoContent)
}

type DeleteMessageRequest struct {
	ID       int64           `json:"id"`
	Message  *models.Message `form:"-" json:"-"`
	AuthInfo *models.ReqInfo `form:"-" json:"-"`
}

func (r *DeleteMessageRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	msg, err := s.Message.GetByID(r.ID)
	if err != nil {
		return err
	}
	if msg == nil {
		return router.NotFoundError()
	}
	r.Message = msg
	r.AuthInfo = router.GetReqInfo(c)
	if !permissions.CanDeleteChat(r.AuthInfo.User, r.Message) {
		return router.PermissionError()
	}
	return nil
}
