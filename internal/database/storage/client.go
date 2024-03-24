package storage

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/veron-baranige/fire-bucket/internal/config"
	"google.golang.org/api/option"
)

var bucket *storage.BucketHandle

func Setup() error {
	parsedCredentials, err := json.Marshal(config.GetFirebaseCredentials())
	if err != nil {
		return err
	}

	opt := option.WithCredentialsJSON(parsedCredentials)
	app, err := firebase.NewApp(
		context.Background(), 
		&firebase.Config{StorageBucket: config.Get(config.FirebaseBucket)}, 
		opt,
	)
	if err != nil {
		return err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return err
	}

	bucket, err = client.DefaultBucket()
	if err != nil {
        return err
    }

	return nil
}
