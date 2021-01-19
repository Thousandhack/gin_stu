package blog

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/blog/get", getHandler)
	e.GET("/blog/comment", commentHandler)
}