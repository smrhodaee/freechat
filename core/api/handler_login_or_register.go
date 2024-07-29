package api

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/router"
	"app/store"
	"app/validations"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type LoginOrRegisterHandler struct {
	s *store.Store
}

func (h *LoginOrRegisterHandler) V1(c *fiber.Ctx) error {
	var req RegisterUserRequest
	if err := router.GetRequest(&req, h.s, c); err != nil {
		return err
	}
	var msg string
	if req.User == nil {
		req.User = &models.User{
			Username: req.Username,
			Password: req.Password,
			Role:     models.UserRoleNormal,
			IsActive: true,
		}
		if err := h.s.User.Create(req.User); err != nil {
			return err
		}
		msg = constants.RegisterSuccessfull
	} else {
		msg = constants.LoginSuccessfull
	}
	token, err := h.s.Token.GenerateRandom(req.User.Username, configs.TokenLength, configs.TokenExpire)
	if err != nil {
		return err
	}
	return c.JSON(router.NewSuccessGlobalResponse(msg, &LoginRegisterResponse{
		Token: token.Value,
	}))
}

type RegisterUserRequest struct {
	Username string       `json:"username" validate:"username"`
	Password string       `json:"password" validate:"password"`
	User     *models.User `form:"-" json:"-"`
}

func (r *RegisterUserRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error {
	r.Username = strings.ToLower(r.Username)
	if !validations.IsValidUsername(r.Username) {
		return router.BadRequestError(constants.InvalidData)
	}
	if !validations.IsValidPassword(r.Password) {
		return router.BadRequestError(constants.InvalidData)
	}
	user, err := s.User.GetByUsername(r.Username)
	if err != nil {
		return err
	}
	r.User = user
	if user != nil {
		if !r.User.IsActive {
			return router.BadRequestError(constants.UserNotActive)
		}
		if !r.User.IsMatchedPassword(r.Password) {
			return router.BadRequestError(constants.InvalidPassword)
		}
	}
	return nil
}

type LoginRegisterResponse struct {
	Token string `json:"token"`
}
