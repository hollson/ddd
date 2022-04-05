package application

import "mime/multipart"

// 适配表现层的数据格式(领域模型无法直接暴露给用户)

type FileInfo struct {
	FileName    string `json:"file_name" form:"file_name"`
	FilePath    string `json:"file_path" form:"file_path"`
	ContentType string `json:"content_type" form:"content_type"`
	Size        int    `json:"size" form:"size"`
}

type AddFileForm struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required"`
	Mark string                `json:"mark" form:"mark"`
}
