package controller

import (
	"go-blog/initializer"
	"go-blog/model"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Check user authentication
	user, _ := c.Get("user")

	// Get data from the body of the request
	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Record data
	post := model.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  user.(model.User).ID,
	}

	// Create the post in the database
	result := initializer.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create Post"})
		return
	}

	// Preload the associated user data
	initializer.DB.Preload("User").First(&post, post.ID)

	// Return the created post with the associated user data
	c.JSON(201, gin.H{"post": post})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post

	initializer.DB.Joins("User").Find(&posts)
	c.JSON(200, gin.H{"posts": posts})
}

func GetUserPosts(c *gin.Context) {
	user, _ := c.Get("user")
	var posts []model.Post
	initializer.DB.Where("user_id = ?", user.(model.User).ID).Find(&posts)
	//initializer.DB.Joins("User").Find(&posts)
	c.JSON(200, gin.H{"posts": posts})
}

func GetPostById(c *gin.Context) {

	id := c.Param("id")
	var post model.Post

	// find data with id
	data := initializer.DB.Joins("User").Find(&post, id)

	if post.ID == 0 {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}
	if data.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}
	// return data
	c.JSON(200, gin.H{"post": post})

}

func UpdatePost(c *gin.Context) {
	// Get id of post
	id := c.Param("id")
	user, _ := c.Get("user")

	var post model.Post

	// Declare a new variable for the request body
	var data struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// Bind body of the request
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find post by id in the database
	findPost := initializer.DB.Joins("User").Find(&post, id)

	// Cover error case
	if findPost.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	if post.UserID != user.(model.User).ID {
		c.JSON(401, gin.H{"message": "Forbidden"})
		return
	}
	// Update post
	initializer.DB.Model(&post).Updates(model.Post{Title: data.Title, Content: data.Content})

	// Return the updated post
	c.JSON(200, gin.H{"post": post})
}

func DeletePostById(c *gin.Context) {
	var post model.Post

	id := c.Param("id")
	user, _ := c.Get("user")

	initializer.DB.Joins("User").Find(&post, id)

	if post.UserID != user.(model.User).ID {
		c.JSON(401, gin.H{"message": "Forbidden"})
		return
	}

	deletePost := initializer.DB.Delete(&model.Post{}, id)

	if deletePost.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, gin.H{"message": "Post has been Deleted!"})

}
