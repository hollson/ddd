package file

import (
	"io/ioutil"
	"time"

	"github.com/google/uuid"
	"github.com/hollson/kendo/config"
	"github.com/hollson/kendo/infrastructure"
	"github.com/hollson/kendo/infrastructure/helper"
	"github.com/hollson/kendo/infrastructure/repo"
)

// 仓储（能封装成工作单元更好）
var captchaRepo = infrastructure.RepoFac.CaptchaRepo
var fileRepo = infrastructure.RepoFac.FilesRepo

// 实体
type fileEntity struct {
	*repo.FileInfo
}

func newEntity(fileInfo FileValue) *fileEntity {
	return &fileEntity{
		FileInfo: &fileInfo.FileInfo,
	}
}

func (en *fileEntity) EntityID() uuid.UUID {
	uid, err := uuid.Parse(en.Id)
	if err != nil {
		uid = uuid.Nil
	}
	return uid
}

func (en *fileEntity) AddFile(body []byte) (err error) {
	filePath := config.FileRoot + time.Now().Format("2006/01/02/") + en.Id + en.FileName
	helper.MakesureFileExist(filePath)
	err = ioutil.WriteFile(filePath, body, 0755)
	if err != nil {
		return
	}
	en.FilePath = filePath
	en.Status = 1
	_, err = fileRepo.AddObj(en.FileInfo)
	return
}
