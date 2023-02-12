package configs

import (
	"log"
	"os"
	"strconv"
	"time"
)

type appConfig struct {
	AppHost            string
	MongoURI           string
	DBName             string
	YoutubeAPIKey      string
	SearchCategory     string
	VideoFetchInterval int64
	VideoFetchFrom     time.Time // value should be in RFC339 format
}

// cfg global container, storing application specific configurations
// for minimizing usage of `os.Getenv()`
var cfg *appConfig

// GetConfig returns the application specific configurations.
func GetConfig() appConfig {
	if cfg != nil {
		return *cfg
	}

	cfg = &appConfig{
		AppHost:            "localhost:3000",
		MongoURI:           "mongodb://localhost:27017/video-search",
		DBName:             "video-search",
		SearchCategory:     "cricket",
		VideoFetchInterval: 10, // in seconds
		VideoFetchFrom:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	val := os.Getenv("APP_HOST")
	if val != "" {
		cfg.AppHost = val
	}

	val = os.Getenv("MONGO_URI")
	if val != "" {
		cfg.MongoURI = val
	}

	val = os.Getenv("MONGO_DB_NAME")
	if val != "" {
		cfg.DBName = val
	}

	val = os.Getenv("YOUTUBE_API_KEY")
	if val != "" {
		cfg.YoutubeAPIKey = val
	}

	val = os.Getenv("SEARCH_CATEGORY")
	if val != "" {
		cfg.SearchCategory = val
	}

	val = os.Getenv("VIDEO_FETCH_INTERVAL")
	if val != "" {
		interval, err := strconv.Atoi(val)
		if err == nil {
			cfg.VideoFetchInterval = int64(interval)
		} else {
			log.Println("provided invalid VIDEO_FETCH_INTERVAL time, using default value")
		}
	}

	val = os.Getenv("VIDEO_FETCH_FROM")
	if val != "" {
		from, err := time.Parse(time.RFC3339, val)
		if err == nil {
			cfg.VideoFetchFrom = from
		} else {
			log.Println("provided invalid VIDEO_FETCH_FROM time, using default value")
		}
	}

	// de-referencing the configs to not alter the configs values
	return *cfg
}
