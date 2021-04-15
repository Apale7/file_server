package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	addr string
)

func init() {
	addr = os.Getenv("file_server_addr")
	if addr == "" {
		addr = ":7777"
	}
}

func main() {
	router := gin.Default()
	defer router.Run(addr)
	register(router)
}
