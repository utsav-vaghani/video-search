package repository

import (
	"context"

	"github.com/utsav-vaghani/video-search/src/domains/video/model"
)

// Video defines repository layer for domain video
type Video interface {
	InsertMany(ctx context.Context, video []model.Video) ([]model.Video, error)
	Find(ctx context.Context, page, limit int) ([]model.Video, error)
	FindByTitle(ctx context.Context, title string) ([]model.Video, error)
	FindByDescription(ctx context.Context, description string) ([]model.Video, error)
}
