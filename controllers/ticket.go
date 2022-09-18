package controllers

import (
	"github.com/KhaleghiDev/jwt-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CreateInput struct {
	Title       string `gorm:"size:150;not null;unique" json:"title"`
	Description string `gorm:"size:150;not null;unique" json:"description"`
	Userid      uint   `gorm:"size:255;not null;" json:"userid"`
}
type FilterInput struct {
	Type   uint `gorm:"not null" json:"type"`
	Status uint `gorm:"not null" json:"status"`
}

func TicketAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ticket []models.Ticket
	db.Find(&ticket)
	c.JSON(http.StatusOK, gin.H{"data": ticket})
	//var input FilterInput
	//t := models.Ticket{}
	//t.Type = input.Type
	//t.Type = input.Status
	//t.Search(t.Type, t.Status)
	//c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi all"})
}

func TicketShow(c *gin.Context) {
	types := c.Query("type")
	search := c.Query("search")
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": search})
	if types != "" {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": types})
	} else if types == "" {
		massages := " not fund"
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": massages})
	}
}

func TicketCreate(c *gin.Context) {

	var input CreateInput
	task := models.Ticket{Title: input.Title, Description: input.Description, UserId: input.Userid}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
}
func TicketDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi create"})
}
func TicketFilter(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi filter"})
}
