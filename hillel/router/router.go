package router

//go get -u github.com/gin-gonic/gin
import (
	"github.com/gin-gonic/gin"
)

// New ...
func New() *gin.Engine {
	r := gin.Default()

	return r
}
