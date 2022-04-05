package infrastructure

import (
	"github.com/hollson/kendo/infrastructure/db/mysql"
	"github.com/hollson/kendo/infrastructure/db/redis"
	repo2 "github.com/hollson/kendo/infrastructure/repo"
)

func init() {
	RepoFac.CaptchaRepo = redis.NewcaptchaRepo()
	RepoFac.FilesRepo = mysql.NewFileRepo()
}

var (
	RepoFac             = &RepoFactory{}
	// Empty   interface{} = struct{}{}
)

type RepoFactory struct {
	CaptchaRepo repo2.CaptchaRepo
	FilesRepo   repo2.FileRepo // 文件仓储」，操作PO(persistent object)持久化对象
}
