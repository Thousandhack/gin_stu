package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			//fmt.Println("文件太大了")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "文件太大了",
			})
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "只允许上传png图片",
			})
			return
		}
		c.SaveUploadedFile(headers, "upload_file/"+headers.Filename)
		c.JSON(http.StatusOK, headers.Filename)
	})
	r.Run()
}

/*
	上传文件过大：
		{
    "message": "文件太大了"
		}
	上传文件不符合文件格式：
		{
    "message": "只允许上传png图片"
		}

*/