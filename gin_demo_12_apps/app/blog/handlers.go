package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello get",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello comment",
	})
}
