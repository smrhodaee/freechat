package api

import (
	"app/configs"
	"app/services/ratelimit"
	"app/store"
	"app/tools"
	"bytes"
	"image/png"

	"github.com/gofiber/fiber/v2"
)

type ChaptaHandler struct {
	s           *store.Store
	rateService *ratelimit.Service
}

func (h *ChaptaHandler) V1(c *fiber.Ctx) error {
	chapta, err := h.s.Chapta.GenerateRandom(configs.ChaptaLength, configs.ChaptaExpire)
	if err != nil {
		return err
	}
	img, err := tools.NewChaptaImage(chapta.Code)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return err
	}
	buf.WriteString("CHAPTA_UUID:" + chapta.UUID)
	c.Set("Content-Type", "image/png")
	return c.SendStream(buf)
}
