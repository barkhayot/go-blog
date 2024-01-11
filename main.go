package main

import (
	"go-blog/controller"
	"go-blog/initializer"
	"go-blog/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.ConnectDB()
	initializer.LoadEnvVariables()
	initializer.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.POST("/logout", middleware.RequireAuth, controller.Logout)
	r.GET("/secure", middleware.RequireAuth, controller.SecureEndpoint)

	// post related endpoints
	r.POST("/posts", middleware.RequireAuth, controller.CreatePost)
	r.GET("/posts", middleware.RequireAuth, controller.GetPosts)
	r.GET("/posts/me", middleware.RequireAuth, controller.GetUserPosts)
	r.GET("/posts/:id", middleware.RequireAuth, controller.GetPostById)
	r.PUT("/posts/:id", middleware.RequireAuth, controller.UpdatePost)
	r.DELETE("/posts/:id", middleware.RequireAuth, controller.DeletePostById)

	r.GET("/check", middleware.RequireAuth, controller.CheckData)

	r.Run()
}
