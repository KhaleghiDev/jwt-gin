package controllers
import(
	"net/http"

	"github.com/gin-gonic/gin"
)
func TicketAll(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"success","data":"hi all"})
}
func TicketShow(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message":"success","data":"hi show"})
}