package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/google/uuid"
)

const (
	MaxFileSize = 10 * 1024 * 1024 // 10MB
	MaxImageSize = 5 * 1024 * 1024  // 5MB
)

var (
	AllowedImageTypes = map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/webp": true,
	}
)

type LocalStorage struct {
	uploadDir string
	baseURL   string
}

func NewLocalStorage(uploadDir, baseURL string) *LocalStorage {
	// Create upload directory if it doesn't exist
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(filepath.Join(uploadDir, "avatars"), 0755)
	os.MkdirAll(filepath.Join(uploadDir, "portfolio"), 0755)
	
	return &LocalStorage{
		uploadDir: uploadDir,
		baseURL:   baseURL,
	}
}

func (s *LocalStorage) SaveAvatar(file *multipart.FileHeader) (string, error) {
	// Validate file type
	if !AllowedImageTypes[file.Header.Get("Content-Type")] {
		return "", fmt.Errorf("invalid file type: %s", file.Header.Get("Content-Type"))
	}

	// Validate file size
	if file.Size > MaxImageSize {
		return "", fmt.Errorf("file too large: %d bytes (max %d)", file.Size, MaxImageSize)
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filepath := filepath.Join(s.uploadDir, "avatars", filename)

	// Save file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return URL
	url := fmt.Sprintf("%s/uploads/avatars/%s", s.baseURL, filename)
	return url, nil
}

func (s *LocalStorage) SavePortfolioImage(file *multipart.FileHeader) (string, error) {
	// Same validation as avatar
	if !AllowedImageTypes[file.Header.Get("Content-Type")] {
		return "", fmt.Errorf("invalid file type: %s", file.Header.Get("Content-Type"))
	}

	if file.Size > MaxImageSize {
		return "", fmt.Errorf("file too large: %d bytes (max %d)", file.Size, MaxImageSize)
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filepath := filepath.Join(s.uploadDir, "portfolio", filename)

	// Save file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return URL
	url := fmt.Sprintf("%s/uploads/portfolio/%s", s.baseURL, filename)
	return url, nil
}

func (s *LocalStorage) DeleteFile(url string) error {
	// Extract filename from URL
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid URL")
	}
	
	filename := parts[len(parts)-1]
	folder := parts[len(parts)-2] // "avatars" or "portfolio"
	
	filepath := filepath.Join(s.uploadDir, folder, filename)
	
	// Delete file if it exists
	if _, err := os.Stat(filepath); err == nil {
		return os.Remove(filepath)
	}
	
	return nil
}
