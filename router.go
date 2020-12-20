package main

import (
	"FileServer/handlers"
	"FileServer/middleware"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	// r.Static("/assets", "./assets")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.POST("/upload", middleware.Mkdir(), handlers.UploadFile)
	r.GET("/download", handlers.DownloadFile)
	r.GET("/ping", handlers.Ping)
}
