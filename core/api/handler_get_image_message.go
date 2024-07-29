package api

import (
	"app/constants"
	"app/models"
	"app/permissions"
	"app/router"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type GetImageMessageHandler struct {
	s *store.Store
}

func (h *GetImageMessageHandler) Register(r fiber.Router, path string) {
}

func (h *GetImageMessageHandler) V1(c *fiber.Ctx) error {
	var req GetImageMessageRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	file, err := h.s.File.GetByID(req.Message.ValueID)
	if err != nil {
		return err
	}
	return c.SendFile(file.Path)
}

type GetImageMessageRequest struct {
	ID       int64           `json:"id"`
	Message  *models.Message `form:"-" json:"-"`
	AuthInfo *models.ReqInfo `form:"-" json:"-"`
}

func (r *GetImageMessageRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	msg, err := s.Message.GetByID(r.ID)
	if err != nil {
		return err
	}
	if msg == nil {
		return router.NotFoundError()
	}
	r.Message = msg
	r.AuthInfo = router.GetReqInfo(c)
	if r.Message.Type != models.MessageTypeImage {
		return router.BadRequestError(constants.NotSupported)
	}
	if !permissions.CanReadFromRoom(r.AuthInfo.User, r.Message.Room) {
		return router.PermissionError()
	}
	return nil
}
