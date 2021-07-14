package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Unit struct {
	Unit string `json:"unit"`
	Long string `json:"long"`
}

type Units []struct {
	ID   int64  `json:"id"`
	Unit string `json:"unit"`
	Long string `json:"long"`
}

func (u Units) contains(id int64) bool {
	for _, n := range u {
		if id == n.ID {
			return true
		}
	}
	return false
}

func unitsPUT(c *gin.Context) {
	var unit Unit
	err := c.BindJSON(&unit)
	if err != nil {
		log.Fatal(err)
	}

	if unit.Unit == "" {
		c.String(http.StatusBadRequest, "Unit fehlt")
	} else {
		id := dbUnitsPUT(unit)
		c.JSON(http.StatusCreated, gin.H{
			"id":   id,
			"unit": unit.Unit,
			"long": unit.Long,
		})
	}
}

func unitsGET(c *gin.Context) {
	res := dbUnitsGET()
	c.JSON(http.StatusOK, res)
}

func unitsPATCH(c *gin.Context) {
	var unit Unit
	err := c.BindJSON(&unit)
	if err != nil {
		log.Fatalf("unitsPATCH BindJSON: %v", err)
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		status := dbUnitsPATCH(id, unit)
		c.Status(status)
	}
}

func unitsDELETE(c *gin.Context) {
	articles := dbArticlesGET()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	switch {
	case err != nil:
		c.Status(http.StatusBadRequest)
	case articles.containsUnit(id):
		c.Status(http.StatusConflict)
	default:
		status := dbDeleteByID("units", id)
		c.Status(status)
	}
}
