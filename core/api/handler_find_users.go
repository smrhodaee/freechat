package api

import (
	"app/constants"
	"app/models"
	"app/router"
	"app/store"
	"app/validations"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type FindUsersRequest struct {
	Username string       `json:"username"`
	AuthUser *models.User `form:"-"`
}

func (r *FindUsersRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	if !validations.IsValidUsername(r.Username) {
		return router.BadRequestError(constants.InvalidData)
	}
	r.AuthUser = router.GetReqInfo(c).User
	r.Username = strings.ToLower(r.Username)
	return nil
}

type FindUsersHandler struct {
	s *store.Store
}

func (h *FindUsersHandler) V1(c *fiber.Ctx) error {
	var req FindUsersRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	users, err := h.s.User.SearchActives(req.Username, req.AuthUser.Username)
	if err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(
		constants.GetObjectsSuccessfully,
		router.Collection(UserResponse{}, users),
	))
}
