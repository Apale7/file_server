package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	defer router.Run(":7777")
	register(router)
}
