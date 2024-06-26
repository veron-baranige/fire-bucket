package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/fire-bucket/internal/config"
	db "github.com/veron-baranige/fire-bucket/internal/database"
	"github.com/veron-baranige/fire-bucket/internal/storage"
)

type (
	FileUploadResponse struct {
		FileId string `json:"fileId"`
	}
)

const (
	maxFileSizeMb = 50
)

// @Summary Upload New File
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Success 201 {object} FileUploadResponse
// @Failure 400
// @Failure 500
// @Router /api/files [post]
// @Tags Files
func SaveFile(c echo.Context) error {
	slog.Info("Received file upload request")

	file, err := c.FormFile("file")
	if err != nil || file == nil {
		slog.Error("Failed to retrieve multipart file header", "err", err, "file", file)
		return echo.ErrBadRequest
	}

	if bytesToMB(file.Size) > float64(maxFileSizeMb) {
		slog.Error("Exceeds max file size", "file.Size", file.Size)
		return echo.ErrBadRequest
	}

	src, err := file.Open()
	if err != nil {
		slog.Error("Failed to extract multipart file", "err", err, "file", file)
		return err
	}
	defer src.Close()

	fileBuff := make([]byte, file.Size)
	if _, err := src.Read(fileBuff); err != nil {
		slog.Error("Failed to read file bytes", "err", err, "file", file)
		return err
	}

	fileId := uuid.NewString()
	fileMimeType := http.DetectContentType(fileBuff)
	filePath := config.Get(config.FileUploadDir) + "/" + fileId

	ctx := context.Background()
	if err := storage.Upload(ctx, storage.FileUpload{
		Path: filePath,
        Content: fileBuff,
		MimeType: fileMimeType,
	}); err!= nil {
		slog.Error("Failed to upload the file to storage bucket", "err", err)
		return err
    }
	slog.Info("Successfully uploaded the file to storage bucket", "path", filePath)

	_, err = db.Client.CreateFile(context.Background(), db.CreateFileParams{
		ID: fileId,
		Name: file.Filename,
		FilePath: filePath,
		Type: fileMimeType,
	})
	if err!= nil {
        slog.Error("Failed to create file record in database", "err", err)
        return echo.ErrInternalServerError
    }
	slog.Info("Successfully saved the file in database", "id", fileId)

    return c.JSON(http.StatusCreated, FileUploadResponse{FileId: fileId})
}

func bytesToMB(bytes int64) float64 {
	return float64(bytes) / 1048576
}

// @Summary Retrieve Uploaded File
// @ID getResourceByID
// @Success 307 
// @Failure 404
// @Failure 500
// @Router /api/files/{id} [get]
// @Tags Files
func GetFile(c echo.Context) error {
	slog.Info("Received file get request")

	id := c.Param("id")

	file, err := db.Client.GetFile(context.Background(), id)
	if err != nil {
		slog.Error("Failed to get file from database", "err", err, "id", id)
		if errors.Is(err, sql.ErrNoRows) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	slog.Error("Retrieved file record from database", "id", id)

	url, err := storage.GetSignedUrl(file.FilePath)
	if err!= nil {
        slog.Error("Failed to get signed url", "err", err, "file.FilePath", file.FilePath)
        return echo.ErrInternalServerError
    }
	slog.Info("Retrieved signed url")

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// @Summary Delete Uploaded File
// @ID getResourceByID
// @Success 204
// @Failure 404
// @Failure 500
// @Router /api/files/{id} [delete]
// @Tags Files
func DeleteFile(c echo.Context) error {
	slog.Info("Received file delete request")

	id := c.Param("id")

	file, err := db.Client.GetFile(context.Background(), id)
	if err != nil {
		slog.Error("Failed to get file from database", "err", err, "id", id)
		if errors.Is(err, sql.ErrNoRows) {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	slog.Info("Retrieved file record from database", "id", id)

	if err := storage.RemoveFromBucket(context.Background(), file.FilePath); err!= nil {
        slog.Error("Failed to remove file from storage bucket", "err", err)
        return echo.ErrInternalServerError
    }
	slog.Info("Removed file from storage bucket")

	if err := db.Client.DeleteFile(context.Background(), id); err!= nil {
        slog.Error("Failed to delete file record from database", "err", err)
        return echo.ErrInternalServerError
    }
	slog.Info("Removed file record from database")

	return c.JSON(http.StatusNoContent, nil)
}
