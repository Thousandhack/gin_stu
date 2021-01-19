package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello checkout",
	})
}

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello goods",
	})
}
