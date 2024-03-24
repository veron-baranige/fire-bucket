package storage

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/veron-baranige/fire-bucket/internal/config"
	"google.golang.org/api/option"
)

type (
	FileUpload struct {
		Content  []byte
		Path     string
		MimeType string
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
	writer := bucket.Object(file.Path).NewWriter(ctx)
	writer.Metadata = map[string]string{
		"Content-Type": file.MimeType,
	}

	if _, err := writer.Write(file.Content); err != nil {
		return err
	}

	return writer.Close()
}

func GetSignedUrl(filePath string) (string, error) {
	return bucket.SignedURL(filePath, &storage.SignedURLOptions{
		Method:  "GET",
		Expires: time.Now().Add(signedUrlExp),
	})
}
