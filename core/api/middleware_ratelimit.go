package api

import (
	"app/constants"
	"app/router"
	"app/services/ratelimit"

	"github.com/gofiber/fiber/v2"
)

type RatelimitMiddleware struct {
	rateService *ratelimit.Service
}

func (r *RatelimitMiddleware) ByIP(rule *ratelimit.Rule) fiber.Handler {
	return func(c *fiber.Ctx) error {
		isExceeded, err := r.rateService.IsExceeded(rule, c.IP())
		if err != nil {
			return err
		}
		if isExceeded {
			return router.BadRequestError(constants.RatelimitError)
		}
		if err := c.Next(); err != nil {
			return err
		}
		status := c.Response().StatusCode()
		if status >= 200 && status < 300 {
			if err = r.rateService.Incr(rule, c.IP()); err != nil {
				return err
			}
		}
		return nil
	}
}
