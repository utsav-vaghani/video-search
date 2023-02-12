package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/utsav-vaghani/video-search/src/common/constants"
	"github.com/utsav-vaghani/video-search/src/domains/video/model"
	"github.com/utsav-vaghani/video-search/src/domains/video/util"
)

type video struct {
	collection *mongo.Collection
}

// New factory function for Video
func New(db *mongo.Database) Video {
	collection := db.Collection(constants.VideosCollection)

	repo := &video{
		collection: collection,
	}

	repo.createIndices()

	return repo
}

func (v *video) InsertMany(ctx context.Context, videos []model.Video) ([]model.Video, error) {
	var documents []interface{}

	for i := range videos {
		videos[i].CreatedID = time.Now().UTC()
		documents = append(documents, videos[i])
	}

	_, err := v.collection.InsertMany(ctx, documents, options.InsertMany().SetBypassDocumentValidation(true))
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (v *video) Find(ctx context.Context, skip, limit int) ([]model.Video, error) {
	pipeline := []bson.M{
		{
			"$skip": skip,
		}, {

			"$limit": limit,
		},
	}

	cursor, err := v.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var videos []model.Video

	for cursor.Next(ctx) {
		var video model.Video

		err = bson.Unmarshal(cursor.Current, &video)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}
	defer cursor.Close(ctx)

	if videos == nil {
		return nil, util.ErrNoVideosFound
	}

	return videos, nil
}

func (v *video) FindByTitle(ctx context.Context, title string) ([]model.Video, error) {
	return v.findByKey(ctx, bson.M{constants.Title: title})
}

func (v *video) FindByDescription(ctx context.Context, description string) ([]model.Video, error) {
	return v.findByKey(ctx, bson.M{constants.Description: description})
}

func (v *video) findByKey(ctx context.Context, key bson.M) ([]model.Video, error) {
	cursor, err := v.collection.Find(ctx, key)
	if err != nil {
		return nil, err
	}

	var videos []model.Video

	for cursor.Next(ctx) {
		var video model.Video

		err = bson.Unmarshal(cursor.Current, &video)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}
	defer cursor.Close(ctx)

	if videos == nil {
		return nil, util.ErrNoVideosFound
	}

	return videos, nil
}

func (v *video) createIndices() {
	indices := []mongo.IndexModel{
		{
			Keys:    bson.M{constants.Title: 1},
			Options: nil,
		},
		{
			Keys:    bson.M{constants.Description: 1},
			Options: nil,
		},
	}

	indexNames, err := v.collection.Indexes().CreateMany(context.Background(), indices)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Indexed: %v\n", indexNames)
}
