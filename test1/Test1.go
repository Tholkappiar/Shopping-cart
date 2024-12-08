package test1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Test1(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "from test 1"})
}
