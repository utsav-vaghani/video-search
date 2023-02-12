package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/utsav-vaghani/video-search/src/common/configs"
	"github.com/utsav-vaghani/video-search/src/core/db/mongodb"
	"github.com/utsav-vaghani/video-search/src/domains/video/controller"
	youtubeService "github.com/utsav-vaghani/video-search/src/domains/video/external/youtube/service"
	"github.com/utsav-vaghani/video-search/src/domains/video/repository"
	"github.com/utsav-vaghani/video-search/src/domains/video/service"
)

func main() {
	app := fiber.New()

	db, err := mongodb.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	cfg := configs.GetConfig()

	videoRepository := repository.New(db)
	youtube := youtubeService.New(cfg.YoutubeAPIKey)
	videoService := service.New(videoRepository, youtube)
	videoController := controller.New(videoService)

	app.Get("/videos", videoController.Get)

	log.Fatal(app.Listen(cfg.AppHost))
}
