package service

import (
	"context"

	"github.com/utsav-vaghani/video-search/src/domains/video/model"
)

// Video defines service layer for domain video
type Video interface {
	GetAll(ctx context.Context, page, limit int) ([]model.Video, error)
	GetByTitle(ctx context.Context, title string) ([]model.Video, error)
	GetByDescription(ctx context.Context, description string) ([]model.Video, error)
}
