package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	
	router.GET("/api/health", healthGET)

	router.GET("/api/boxes", boxesGET)
	router.PUT("/api/boxes", boxesPUT)
	router.PATCH("/api/boxes/:id", boxesPATCH)
	router.DELETE("/api/boxes/:id", boxesDELETE)
	
	router.GET("/api/units", unitsGET) 
	router.PUT("/api/units", unitsPUT) 
	router.PATCH("/api/units/:id", unitsPATCH)
	router.DELETE("/api/units/:id", unitsDELETE)
	
	return router
}
