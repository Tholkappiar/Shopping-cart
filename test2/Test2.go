package test2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func Test2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "from test 2"})
}

