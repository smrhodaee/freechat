package router

import (
	"app/configs"
	"app/constants"
	"app/models"
	"app/store"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// type Middleware interface {
// 	Handler(c *fiber.Ctx) error
// }

// type Handler interface {
// 	V1(c *fiber.Ctx) error
// }

type Request interface {
	ValidateAndNormalize(s *store.Store, c *fiber.Ctx) error
}

type Response[R any] interface {
	Clone(data any) R
}

func Collection[T any, R any](res Response[R], items []T) []R {
	var ret []R
	for _, item := range items {
		ret = append(ret, res.Clone(item))
	}
	return ret
}

func GetRequest(req Request, s *store.Store, c *fiber.Ctx) error {
	if err := c.BodyParser(req); err != nil {
		return BadRequestError(constants.InvalidRequest)
	}
	if err := req.ValidateAndNormalize(s, c); err != nil {
		return err
	}
	return nil
}

func GetLocal[T any](c *fiber.Ctx, name string) *T {
	ret, ok := c.Locals(name).(*T)
	if ok {
		return ret
	}
	return nil
}

func GetReqInfo(c *fiber.Ctx) *models.ReqInfo {
	return GetLocal[models.ReqInfo](c, configs.LocalReqInfo)
}

func GetReqInfoWs(c *websocket.Conn) *models.ReqInfo {
	req, ok := c.Locals(configs.LocalReqInfo).(*models.ReqInfo)
	if ok {
		return req
	}
	return nil
}
