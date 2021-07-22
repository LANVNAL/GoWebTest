package main

import (
	"GoWebTest/tools"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

type LoginForm struct {
	User string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	//登录
	router.POST("/login", func(c *gin.Context){
	var form LoginForm
	if c.ShouldBind(&form) == nil{
		if form.User == "admin" && form.Password == "password"{
			c.SetCookie("user_cookie", "admin", 9999, "/", "localhost", false, true)
			c.JSON(200, gin.H{
				"msg": "welcome admin.Login Success!",
			})
		}else {
			c.JSON(401, gin.H{
				"msg": "unauthorized, Try Again!",
			})
		}
	}else {
		c.JSON(401, gin.H{
			"msg": "missing params",
		})
	}
	})
	
	//验证登录态
	router.GET("/checklogin", func(c *gin.Context) {
		cookie, err := c.Cookie("user_cookie")

		if err == nil{
			//fmt.Println(cookie)
			c.JSON(200, gin.H{
				"cookie": cookie,
			})
		}
	})

	router.POST("/fileupload", func(context *gin.Context) {
		//单文件上传
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": "get form err: " + err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		filename := filepath.Base(file.Filename)
		var dst = "./fileupload/" + filename
		if err := context.SaveUploadedFile(file, dst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": "upload file err: " + err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": "File " + file.Filename + " uploaded successfully.",
		})
	})
	
	
	//获取全部文件名称
	router.GET("/getall", func(context *gin.Context) {
		fileList, _ := tools.ListAll("./fileupload")
		context.JSON(http.StatusOK, gin.H{
			"msg": fileList,
		})
	})

	router.Run(":8888")
}
