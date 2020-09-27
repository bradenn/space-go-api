package Router

import (
	"github.com/gin-gonic/gin"
	"space-api/Controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/systems/:id", Controller.GetSystem)
		v1.GET("/systems", Controller.GetAllSystems)
		v1.GET("/stars/:id", Controller.GetStar)
		v1.GET("/stars/", Controller.GetAllStars)
		v1.GET("/planets/:id", Controller.GetPlanet)
		v1.GET("/planets/", Controller.GetAllPlanets)
		v1.GET("/moons/", Controller.GetAllMoons)
		v1.GET("/moons/:id", Controller.GetMoon)
	}
	return r
}
