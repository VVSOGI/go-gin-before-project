package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vvsogi/goginstudy/initializers"
	"github.com/vvsogi/goginstudy/models"
)

var body struct {
	Body 	string
	Title 	string
}

// function that create a post
func PostsCreate (c *gin.Context) {
	// Get data off req body
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: "benny", Body: "is soooo good"}
	result := initializers.Database.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	}) 
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.Database.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
