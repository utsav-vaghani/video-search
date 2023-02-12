package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/utsav-vaghani/video-search/src/domains/video/util"
)

func convertToFiberError(err error) error {
	if err == nil {
		return err
	}

	if _, ok := err.(*fiber.Error); ok {
		return err
	}

	var code = fiber.StatusInternalServerError

	if err == util.ErrNoVideosFound {
		code = fiber.StatusInternalServerError
	}

	return fiber.NewError(code, err.Error())
}
