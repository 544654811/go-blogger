package main

import (
	"blogger/controller"
	"blogger/dal/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置gin路由
	router := gin.Default()

	// 初始化mysql数据库配置
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.InitDb(dns)
	if err != nil {
		fmt.Println("初始化数据库错误..", err)
		return
	}

	// gin 初始化静态资源
	router.Static("/static/", "./static")
	// gin html模板
	router.LoadHTMLGlob("views/*")

	// 配置路由
	router.GET("/", controller.ToIndex)
	router.GET("/article/detail", controller.ToArticleDetail)
	router.GET("/article/new", controller.ToPostArticle)

	router.POST("/article/submit/", controller.PostArticleHandle)

	router.POST("/comment/submit/", controller.PostCommentHandle)

	router.GET("/leave/new/", controller.ToLeave)
	router.POST("/leave/submit/", controller.PostLeaveHandle)

	router.GET("/category/", controller.ToCategory)
	router.Run(":9090")
}
