package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadShop(e *gin.Engine)  {
	e.GET("/hello", helloHandler)
	e.GET("/goods", goodsHandler)

}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}


func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello goods",
	})
}