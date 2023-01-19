package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vvsogi/goginstudy/controllers"
	"github.com/vvsogi/goginstudy/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
