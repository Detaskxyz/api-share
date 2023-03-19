package models

import (
	"time"

	"gorm.io/gorm"
)

type FileUpload struct {
	gorm.Model
	Filename   string
	URL        string
	Size       int64
	Key        string
	UploadTime time.Time
	IP         string
	UID        uint
}

func (table FileUpload) TableName() string {
	return "file_upload"
}
