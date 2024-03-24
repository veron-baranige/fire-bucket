package storage

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/veron-baranige/fire-bucket/internal/config"
	"google.golang.org/api/option"
)

type (
	FileUpload struct {
		content []byte
		path    string
	}
)

const (
	signedUrlExp = time.Minute * 2
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

func Upload(ctx context.Context, file FileUpload) error {
	writer := bucket.Object(file.path).NewWriter(ctx)
	writer.Metadata = map[string]string{
		"Content-Type": http.DetectContentType(file.content),
	}

	if _, err := writer.Write(file.content); err != nil {
		return err
	}
	return nil
}

func GetSignedUrl(filePath string) (string, error) {
	url, err := bucket.SignedURL(filePath, &storage.SignedURLOptions{
		Method: "GET",
		Expires: time.Now().Add(signedUrlExp),
	})
	if err!= nil {
        return "", err
    }
	return url, nil
}
