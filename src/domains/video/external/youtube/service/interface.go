package service

import (
	"time"

	"github.com/utsav-vaghani/video-search/src/domains/video/external/youtube/model"
)

// Youtube defining methods for youtube service layer
type Youtube interface {
	List(category, resourceType, order, nextPageToken string, publishedAfter *time.Time) (*model.VideoResponse, error)
}
