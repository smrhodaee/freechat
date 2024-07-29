package configs

import (
	"fmt"
	"mime/multipart"
	"time"
)

func UploadFilePath(file *multipart.FileHeader) string {
	return fmt.Sprintf("./uploads/%d-%s", time.Now().UnixMicro(), file.Filename)
}
