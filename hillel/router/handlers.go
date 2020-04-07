package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func collectHandler() {

}

func resultHandler() {

}
