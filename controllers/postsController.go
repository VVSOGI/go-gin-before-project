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
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.Database.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
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

func PostsShow(c *gin.Context) {
	// Get the posts
	id := c.Param("id")
	var posts []models.Post
	initializers.Database.First(&posts, id)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	DB := initializers.Database
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	DB.First(&post, id)

	// Update it
	DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	DB := initializers.Database

	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
