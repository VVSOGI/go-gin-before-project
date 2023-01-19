package main

import (
	"github.com/vvsogi/goginstudy/initializers"
	"github.com/vvsogi/goginstudy/models"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
  }

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	DB := initializers.ConnectToDB()
	DB.AutoMigrate(&models.Post{})
}
