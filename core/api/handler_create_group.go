package api

import (
	"app/constants"
	"app/models"
	"app/router"
	"app/store"
	"app/tools"
	"app/validations"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CreateGroupRequest struct {
	Name      string          `json:"name"`
	Title     string          `json:"title"`
	Usernames []string        `json:"usernames"`
	Room      *models.Room    `form:"-" json:"-"`
	Users     []models.User   `form:"-" json:"-"`
	AuthInfo  *models.ReqInfo `form:"-" json:"-"`
}

func (r *CreateGroupRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	if !validations.IsValidGroupName(r.Name) {
		return router.BadRequestError(constants.InvalidData)
	}
	if !validations.IsValidRoomTitle(r.Title) {
		return router.BadRequestError(constants.InvalidData)
	}
	r.Name = strings.ToLower(r.Name)
	r.Usernames = tools.UniqueString(r.Usernames)
	if len(r.Usernames) == 0 {
		return router.BadRequestError(constants.MemberError)
	}
	if len(r.Usernames) > 10 {
		return router.BadRequestError(constants.RatelimitError)
	}
	room, err := s.Room.GetByName(r.Name)
	if err != nil {
		return err
	}
	if room != nil {
		return router.BadRequestError(constants.RoomExists)
	}
	users, err := s.User.GetActivesByUsername(r.Usernames)
	if err != nil {
		return err
	}
	var reqInfo = router.GetReqInfo(c)
	if tools.FindElement(r.Usernames, func(value *string) bool {
		return strings.Compare(*value, reqInfo.Username) == 0
	}) != nil {
		return router.BadRequestError(constants.MemberError)
	}
	if users == nil || len(users) != len(r.Usernames) {
		return router.PermissionError()
	}
	r.Room = room
	r.Users = users
	r.AuthInfo = router.GetReqInfo(c)
	return nil
}

type CreateGroupHandler struct {
	s *store.Store
}

func (h *CreateGroupHandler) V1(c *fiber.Ctx) error {
	var req CreateGroupRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	if err := h.s.Room.CreateGroup(req.Name, req.Title, req.AuthInfo.User, req.Users); err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(constants.RoomCreatedSuccessfuly))
}
