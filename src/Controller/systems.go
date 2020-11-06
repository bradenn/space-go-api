package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"space-api/src/Models"
)

func GetSystem(c *gin.Context) {
	id := c.Params.ByName("id")
	var system Models.System
	err := Models.FindSystem(&system, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, system)
	}

}

func GetAllSystems(c *gin.Context) {
	var systems []Models.System
	err := Models.FindAllSystems(&systems)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, systems)
	}

}