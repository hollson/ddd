package application

import (
	"context"
	"io/ioutil"

	"github.com/hollson/kendo/domain/file"
	"github.com/hollson/kendo/infrastructure"
	"github.com/hollson/kendo/infrastructure/bus"
	"github.com/hollson/kendo/infrastructure/ddd"
	eh "github.com/looplab/eventhorizon"
)

// GetFileById 「从数据仓储」获取持久化数据(PO)，转换为视图对象(VO)并暴露给表现层
func GetFileById(id string) (vo FileInfo, has bool, err error) {
	obj, has, err := infrastructure.RepoFac.FilesRepo.GetById(id)
	if err != nil || !has {
		return
	}
	vo.ContentType = obj.ContentType
	vo.FileName = obj.FileName
	vo.FilePath = obj.FilePath
	vo.Size = obj.Size
	return
}

// AddFile 操作领域层
func AddFile(vm AddFileForm) (fileId string, err error) {
	f, err := vm.File.Open()
	if err != nil {
		return
	}
	defer f.Close()

	vo := file.FileValue{}
	vo.FileBody, err = ioutil.ReadAll(f)
	if err != nil {
		return
	}

	vo.ContentType = vm.File.Header.Get("Content-Type")
	vo.Size = int(vm.File.Size)
	vo.FileName = vm.File.Filename
	fileId, err = file.SingleFilesAgg.AddNewFile(vo)

	return
}

// AddFileCommand 框架根据命令类型，创建命令实例
func AddFileCommand(vm AddFileForm) (err error) {
	cmd, err := eh.CreateCommand(file.AddFileCmdType)
	if err != nil {
		return
	}

	// 命令进行校验
	if vldt, ok := cmd.(ddd.Validator); ok {
		err = vldt.Verify()
		if err != nil {
			return
		}
	}

	// 通过命令总线，将命令发布出去，至于谁订阅了该命令则不关心
	if err = bus.HandleCommand(context.Background(), cmd); err != nil {
		return err
	}
	return
}
