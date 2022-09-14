package controllers

import (
	"github.com/KhaleghiDev/jwt-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateInput struct {
	Title       string `gorm:"size:150;not null;unique" json:"title"`
	Description string `gorm:"size:150;not null;unique" json:"description"`
	Userid      uint   `gorm:"size:255;not null;" json:"user_id"`
}

func TicketAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi all"})
}
func TicketShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi show"})
}
func TicketCreate(c *gin.Context) {
	var input CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := models.Ticket{}

	t.Title = input.Title
	t.Description = input.Description
	// user_id, _ := strconv.ParseUint(input.Userid)
	t.UserId = input.Userid

	_, err := t.SaveTicket()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ticke success"})

}
func TicketDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi create"})
}
func TicketFilter(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": "hi filter"})
}
