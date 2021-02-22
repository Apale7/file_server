package handlers

import (
	_ "file_server/dal"
	"file_server/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	Base = "~/storage/"
)

func UploadFile(c *gin.Context) {
	fmt.Println("UploadFile called")
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		utils.RetErr(c, err)
		return
	}
	logrus.WithContext(c).Debugf("Get file: ", file.Filename)
	fmt.Printf("%s", file.Filename)
	if err = c.SaveUploadedFile(file, Base+file.Filename); err != nil {
		utils.RetErr(c, fmt.Errorf("upload file err: %s", err.Error()))
		return
	}
	utils.RetData(c, file.Filename)
}

func DownloadFile(c *gin.Context) {
	filepath := c.Query("filepath")
	if filepath == "" {
		utils.RetErr(c, fmt.Errorf("未输入文件路径"))
		return
	}
	router := strings.Split(filepath, "/")
	filename := router[len(router)-1]
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(Base + filepath)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
