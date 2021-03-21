package handlers

import (
	_ "file_server/dal"
	"file_server/utils"
	"fmt"
	"os"
	"strings"

	"github.com/Apale7/common/constdef"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	Base = "/Users/apale/storage/"
	Host = "127.0.0.1:7777"
)

type uploadParams struct {
	UserID uint32 `json:"user_id" form:"user_id"`
	Type   string `json:"type" form:"type"`
}

func UploadFile(c *gin.Context) {
	fmt.Println("UploadFile called")
	var reqBody uploadParams
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("%+v\n", reqBody)
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		utils.RetErr(c, err)
		return
	}
	path := fmt.Sprintf("%d/%s/", reqBody.UserID, reqBody.Type)
	logrus.Infoln(Base + path)
	err = os.MkdirAll(Base+path, 0777)
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.WithContext(c).Debugf("Get file: ", file.Filename)
	fmt.Printf("%s", file.Filename)
	if err = c.SaveUploadedFile(file, Base+path+file.Filename); err != nil {
		utils.RetErr(c, fmt.Errorf("upload file err: %s", err.Error()))
		return
	}
	utils.RetData(c, fmt.Sprintf("%s/download?filepath=%s", Host, path+file.Filename))
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
	// c.String(200, "asd")
	c.File(Base + filepath)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
