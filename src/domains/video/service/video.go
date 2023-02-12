package service

import (
	"context"
	"log"
	"time"

	"github.com/utsav-vaghani/video-search/src/common/configs"
	youtubeService "github.com/utsav-vaghani/video-search/src/domains/video/external/youtube/service"
	"github.com/utsav-vaghani/video-search/src/domains/video/model"
	"github.com/utsav-vaghani/video-search/src/domains/video/repository"
)

type video struct {
	repo           repository.Video
	youtubeService youtubeService.Youtube
}

// New factory function for Video
func New(repo repository.Video, youtubeService youtubeService.Youtube) Video {
	service := &video{
		repo:           repo,
		youtubeService: youtubeService,
	}

	go service.fetchYoutubeVideos()

	return service
}

func (v *video) GetAll(ctx context.Context, page, limit int) ([]model.Video, error) {
	if limit == 0 {
		limit = 10 // default: 10 videos will be returned
	}

	skip := page * limit

	return v.repo.Find(ctx, skip, limit)
}

func (v *video) GetByTitle(ctx context.Context, title string) ([]model.Video, error) {
	return v.repo.FindByTitle(ctx, title)
}

func (v *video) GetByDescription(ctx context.Context, description string) ([]model.Video, error) {
	return v.repo.FindByDescription(ctx, description)
}

func (v *video) fetchYoutubeVideos() {
	cfg := configs.GetConfig()
	publishedAfter := cfg.VideoFetchFrom

	var nextPageToken string

	for {
		resp, err := v.youtubeService.List(cfg.SearchCategory, "video", "date", nextPageToken, &publishedAfter)
		if err != nil {
			log.Printf("error while fetching videos from youtube, %v\n", err)

			time.Sleep(time.Second * 10)
			continue
		}

		var videos []model.Video

		for _, item := range resp.Items {
			video := model.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: item.Snippet.PublishedAt,
				ThumbnailURLs: model.ThumbnailURLs{
					Default: item.Snippet.Thumbnails.Default.Url,
					Medium:  item.Snippet.Thumbnails.Medium.Url,
					High:    item.Snippet.Thumbnails.High.Url,
				},
			}

			videos = append(videos, video)
		}

		log.Printf("fetched %v videos from youtube\n", len(videos))

		if len(videos) > 0 {
			_, err = v.repo.InsertMany(context.Background(), videos)
			if err != nil {
				log.Printf("error while storing fetched videos, %v\n", err)
			}
		}

		if resp.NextPageToken != "" {
			nextPageToken = resp.NextPageToken
		} else {
			publishedAfter = time.Now()
			nextPageToken = ""
		}

		time.Sleep(time.Second * 10)
	}
}
