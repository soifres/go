package router

import (
	"github.com/gin-gonic/gin"
)

//go get -u github.com/gin-gonic/gin

// New ...
func New() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexHandler)
	r.GET("/search/:category", collectHandler)
	r.GET("/result/:category", resultHandler)
	return r
}
