package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"space-api/src/Models"
)

func GetStar(c *gin.Context) {
	id := c.Params.ByName("id")
	var star Models.Star
	err := Models.FindStar(&star, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, star)
	}

}

func GetAllStars(c *gin.Context) {
	var stars []Models.Star
	err := Models.FindAllStars(&stars)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stars)
	}

}