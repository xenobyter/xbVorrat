package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthGET(c *gin.Context) {
	var res = make(map[string]int)
	res["version"] = 1
	c.JSON(http.StatusOK, res)
}
