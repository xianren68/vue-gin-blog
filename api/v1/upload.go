package v1

import (
	"gin-blog/errormsg"
	"gin-blog/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadApi struct {
}

func (ul *UploadApi) UploadImg(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	// 获取文件长度
	fileSize := fileHeader.Size
	url, code := upload.UploadImg(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errormsg.GetMsg(code),
		"url":    url,
	})
}
