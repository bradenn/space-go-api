package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"space-api/src/Models"
)

func GetMoon(c *gin.Context) {
	id := c.Params.ByName("id")
	var moon Models.Moon
	err := Models.FindMoon(&moon, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, moon)
	}

}

func GetAllMoons(c *gin.Context) {
	var moons []Models.Moon
	err := Models.FindAllMoons(&moons)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, moons)
	}

}