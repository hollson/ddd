// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hollson/kendo/application"
	"github.com/hollson/kendo/infrastructure/errorext"
)

func AddFileHandler(c *gin.Context) {
	var req application.AddFileForm
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}
	if req.File == nil {
		c.Error(errorext.NewCodeError(101, "文件无效", nil))
		return
	}
	fid, err := application.AddFile(req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, fid)
}
