package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hollson/kendo/application"
)

func GetFileHandler(c *gin.Context) {
	rst, has, err := application.GetFileById(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	if !has {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", rst.FileName))
	c.Header("Content-Type", rst.ContentType)
	c.File(rst.FilePath)
}
