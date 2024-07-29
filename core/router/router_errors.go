package router

import (
	"app/constants"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ParseError(err error) *fiber.Error {
	var fErr = fiber.NewError(fiber.StatusInternalServerError, constants.ServerError)
	errors.As(err, &fErr)
	return fErr
}

func BadRequestError(msg string) *fiber.Error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func PermissionError() *fiber.Error {
	return fiber.NewError(fiber.StatusForbidden, constants.PermissionDenied)
}

func NotFoundError() *fiber.Error {
	return fiber.NewError(fiber.StatusNotFound, constants.NotFound)
}
