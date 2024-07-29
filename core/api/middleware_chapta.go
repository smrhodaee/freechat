package api

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/router"
	"app/services/ratelimit"
	"app/store"

	"github.com/gofiber/fiber/v2"
)

type ChaptaRequest struct {
	models.Chapta
}

func (*ChaptaRequest) ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error { return nil }

type ChaptaMiddleware struct {
	s           *store.Store
	rateService *ratelimit.Service
}

func (m *ChaptaMiddleware) Handler(c *fiber.Ctx) error {
	var req ChaptaRequest
	if err := router.GetRequest(&req, m.s, c); err != nil {
		return err
	}
	exceeded, err := m.rateService.IsExceeded(configs.ChaptaRatelimit, c.IP())
	if err != nil {
		return err
	}
	if exceeded {
		return router.BadRequestError(constants.RatelimitError)
	}
	valid, err := m.s.Chapta.IsValid(&req.Chapta)
	if err != nil {
		return err
	}
	if !valid {
		return router.BadRequestError(constants.InvalidChapta)
	}
	if err := m.s.Chapta.Delete(req.UUID); err != nil {
		return err
	}
	return c.Next()
}
