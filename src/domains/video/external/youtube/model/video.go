package model

import "time"

// Video model for youtube data v3 search api
type Video struct {
	Snippet struct {
		PublishedAt time.Time `json:"publishedAt"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Thumbnails  struct {
			Default struct {
				Url string `json:"url"`
			} `json:"default"`
			Medium struct {
				Url string `json:"url"`
			} `json:"medium"`
			High struct {
				Url string `json:"url"`
			} `json:"high"`
		} `json:"thumbnails"`
	} `json:"snippet"`
}

// VideoResponse of youtube search api
type VideoResponse struct {
	NextPageToken string  `json:"nextPageToken"`
	Items         []Video `json:"items"`
}
