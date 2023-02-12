package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/utsav-vaghani/video-search/src/domains/video/external/youtube/model"
)

type youtube struct {
	apiKey string
}

func New(apiKey string) Youtube {
	return &youtube{
		apiKey: apiKey,
	}
}

func (y *youtube) List(category, resourceType, order, nextPageToken string, publishedAfter *time.Time) (*model.VideoResponse, error) {
	if publishedAfter == nil {
		// this can be enhanced by storing/caching somewhere to handle case of instance failure.
		now := time.Now()
		publishedAfter = &now
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&part=snippet&q=%s&type=%s&order=%s&publishedAfter=%s&maxResults=50",
		y.apiKey, category, resourceType, order, publishedAfter.Format(time.RFC3339))

	if nextPageToken != "" {
		url += fmt.Sprintf("&pageToken=%s", nextPageToken)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.VideoResponse

	err = json.Unmarshal(bodyData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
