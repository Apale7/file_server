package middleware

import (
	"file_server/handlers"
	"file_server/utils"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Mkdir 中间件用来判断上传路径中的文件夹是否都存在，不存在则依次创建
func Mkdir() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil || file == nil {
			utils.RetErr(c, err)
			c.Abort()
			return
		}
		path := handlers.Base
		router := strings.Split(file.Filename, "/")
		fmt.Printf("%+v", router)
		for _, dir := range router[0 : len(router)-1] {
			path += dir + "/"
			fmt.Printf("%s", path)
			flag := pathExists(path)
			if !flag {
				os.Mkdir(path, os.FileMode(511))
			}
		}
		c.Next()
	}
}
