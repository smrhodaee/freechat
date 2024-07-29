package router

import (
	"app/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

type GlobalResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessGlobalResponse(msg string, data ...interface{}) *GlobalResponse {
	return &GlobalResponse{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}

func New(log *logrus.Logger, store *store.Store) *fiber.App {
	log.SetReportCaller(false)
	r := fiber.New(
		fiber.Config{
			BodyLimit:             4 * 1024 * 1024,
			DisableStartupMessage: false,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				if err != nil {
					fErr := ParseError(err)
					if fErr.Code == fiber.StatusInternalServerError {
						log.Errorf("%s %s [%s] by (%s)", c.Method(), c.OriginalURL(), err, c.IP())
					}
					return c.Status(fErr.Code).JSON(GlobalResponse{
						Status:  false,
						Message: fErr.Message,
					})
				}
				return nil
			},
		},
	)
	// r.Use(logger.New())
	r.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
		}),
	)
	r.Use(authMiddleware(store))
	return r
}
