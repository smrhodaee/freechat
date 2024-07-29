package models

import (
	"image"
	"os"
)

type File struct {
	ID            int64  `gorm:"primaryKey" json:"id"`
	OwnerUsername string `json:"owner_username"`
	Title         string `json:"title"`
	Path          string `json:"path"`
	Size          int64 `json:"size"`
	MimeType      string `json:"mime_type"`
	IsActive      bool   `json:"is_active"`
}

type ImageInfo struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (f *File) GetImageInfo() (*ImageInfo, error) {
	if f.MimeType == "image/jpeg" || f.MimeType == "image/png" {
		reader, err := os.Open(f.Path)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		im, _, err := image.DecodeConfig(reader)
		if err != nil {
			return nil, err
		}
		return &ImageInfo{
			Width:  im.Width,
			Height: im.Height,
		}, nil
	}
	return nil, nil
}

// func (File) Table() string {
// 	return "file"
// }
