package api

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/router"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type CreateDirectRequest struct {
	Username string          `json:"username"`
	User     *models.User    `form:"-" json:"-"`
	AuthInfo *models.ReqInfo `form:"-" json:"-"`
}

func (r *CreateDirectRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	user, err := s.User.GetByUsername(r.Username)
	if err != nil {
		return err
	}
	if user == nil {
		return router.NotFoundError()
	}
	r.User = user
	r.AuthInfo = router.GetReqInfo(c)
	//TODO: block list
	return nil
}

type CreateDirectHandler struct {
	s *store.Store
}

func (h *CreateDirectHandler) V1(c *fiber.Ctx) error {
	var req CreateDirectRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	roomName := configs.GetDirectRoomName(req.AuthInfo.User, req.User)
	room, err := h.s.Room.GetByName(roomName)
	if err != nil {
		return err
	}
	if room != nil {
		return router.BadRequestError(constants.RoomExists)
	}
	if err := h.s.Room.CreateDirect(roomName, req.AuthInfo.User, req.User); err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(constants.RoomCreatedSuccessfuly))
}
