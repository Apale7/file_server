package middleware

import (
	"FileServer/handlers"
	"FileServer/utils"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

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
			flag, _ := pathExists(path)
			if !flag {
				os.Mkdir(path, os.FileMode(511))
			}
		}
		c.Next()
	}
}
