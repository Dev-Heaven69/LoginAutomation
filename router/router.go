package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seew0/loginAutomation/logic"
	"github.com/seew0/loginAutomation/models"
)

type Router struct {
	Logic logic.Logic
}

func ProvideRouter(l logic.Logic) Router {
	return Router{l}
}

func (r Router) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server healthy"})
}

func (r Router) GetLoginSmartlead(c *gin.Context) {
	var req models.GmailLoginSmartlead
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = r.Logic.GetLoginSmartlead(req, c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Completed Gmail Login and Smartlead account insertion"})
}

func (r Router) UpdateFiveSimOptions(c *gin.Context){
	var req models.FiveSimOptions
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.Logic.UpdateFiveSimOptions(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated fivesim options"})
}