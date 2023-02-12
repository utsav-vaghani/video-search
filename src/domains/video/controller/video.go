package controller

import (
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/utsav-vaghani/video-search/src/common/constants"
	"github.com/utsav-vaghani/video-search/src/domains/video/service"
)

type video struct {
	service service.Video
}

// New factory function for Video
func New(service service.Video) Video {
	return &video{
		service: service,
	}
}

func (v *video) Get(ctx *fiber.Ctx) error {
	title := ctx.Query(constants.Title)
	description := ctx.Query(constants.Description)

	if title != "" {
		videos, err := v.service.GetByTitle(ctx.UserContext(), title)
		if err != nil {
			return convertToFiberError(err)
		}

		return ctx.JSON(videos)
	} else if description != "" {
		videos, err := v.service.GetByDescription(ctx.UserContext(), description)
		if err != nil {
			return convertToFiberError(err)
		}

		return ctx.JSON(videos)
	}

	var (
		page  int
		limit int
	)

	if val := ctx.Query(constants.Page); val != "" {
		pageParam, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		page = pageParam
	}

	if val := ctx.Query(constants.Limit); val != "" {
		limitParam, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		limit = limitParam
	}

	videos, err := v.service.GetAll(ctx.UserContext(), page, limit)
	if err != nil {
		return convertToFiberError(err)
	}

	return ctx.JSON(videos)
}
