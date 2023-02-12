package model

import "time"

// Video defines the fields related to youtube video
type Video struct {
	CreatedID     time.Time     `json:"createdId" bson:"createdId"`
	Title         string        `json:"title" bson:"title"`
	Description   string        `json:"description" bson:"description"`
	PublishedAt   time.Time     `json:"publishedAt" bson:"publishedAt"`
	ThumbnailURLs ThumbnailURLs `json:"thumbnailURLs" bson:"thumbnailURLs"`
}

// ThumbnailURLs contains thumbnail URL with different quality
type ThumbnailURLs struct {
	Default string `json:"default,omitempty" bson:"default,omitempty"`
	Medium  string `json:"medium,omitempty" bson:"medium,omitempty"`
	High    string `json:"high,omitempty" bson:"high,omitempty"`
}
