package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"space-api/src/Models"
)

func GetPlanet(c *gin.Context) {
	id := c.Params.ByName("id")
	var planet Models.Planet
	err := Models.FindPlanet(&planet, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, planet)
	}

}

func GetAllPlanets(c *gin.Context) {
	var planets []Models.Planet
	err := Models.FindAllPlanets(&planets)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, planets)
	}

}