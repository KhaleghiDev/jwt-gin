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

	v1 := r.Group("/api/v1")
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)

	v1admin := r.Group("/api/v1/admin")
	v1admin.Use(middlewares.JwtAuthMiddleware())
	v1admin.GET("/user", controllers.CurrentUser)

	v1admin.Use(middlewares.JwtAuthMiddleware())
	v1admin.GET("ticket/all", controllers.TicketAll)
	v1admin.POST("ticket/create", controllers.TicketCreate)
	v1admin.GET("ticket/filter", controllers.TicketShow)

	r.Run(":8080")

}
