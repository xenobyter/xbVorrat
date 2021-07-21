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

	router.GET("/api/articles", articlesGET)
	router.PUT("/api/articles", articlesPUT)
	router.PATCH("/api/articles/:id", articlesPATCH)
	router.DELETE("/api/articles/:id", articlesDELETE)

	router.GET("/api/stocks", stocksGET)
	router.GET("/api/stocks/rich", stocksRichGET)
	router.PUT("/api/stocks", stocksPUT)
	router.PATCH("/api/stocks/:id", stocksPATCH)
	router.DELETE("/api/stocks/:id", stocksDELETE)

	router.StaticFile("/", "app/dist/index.html")
	router.Static("/css", "app/dist/css")
	router.Static("/fonts", "app/dist/fonts")
	router.Static("/js", "app/dist/js")
	router.Static("/icons", "app/dist/icons")

	return router
}
