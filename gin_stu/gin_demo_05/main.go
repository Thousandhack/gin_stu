package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	//r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {

		// 表单取文件
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传文件出错")
		}
		log.Println(file.Filename)
		// c.JSON(200, gin.H{"message": file.Header.Context})
		// 传到项目目录，名字就用本身
		//c.SaveUploadedFile(file, file.Filename)
		// 传到自己指定的目录下，下面就会上传到src下gin_stu/upload_file/ 的目录
		c.SaveUploadedFile(file, "./upload_file/"+file.Filename)
		//c.String(http.StatusOK, file.Filename)
		c.JSON(http.StatusOK,gin.H{
			"message": "save "+file.Filename,
		})
	})
	r.Run()
}

/*
	成功的返回结果：
{
    "message": "save black.jpg"
}
*/