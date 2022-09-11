package main

import (
	"github.com/KhaleghiDev/jwt-gin/controllers"
	"github.com/KhaleghiDev/jwt-gin/middlewares"
	"github.com/KhaleghiDev/jwt-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api/v1")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	
	ticket := r.Group("/api/v1/admin")
	ticket.Use(middlewares.JwtAuthMiddleware())
	ticket.GET("/user",controllers.CurrentUser)

	r.Run(":8080")

}
