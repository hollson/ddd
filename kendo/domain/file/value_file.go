package file

import (
	"github.com/hollson/kendo/infrastructure/repo"
)

// FileValue 值对象(value object)或(view object)表现层对象
type FileValue struct {
	repo.FileInfo
	FileBody []byte
}
