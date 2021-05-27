package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Unit struct {
	Unit  string `json:"unit"`
	Long string `json:"long"`
}

type Units []struct {
	ID    int64  `json:"id"`
	Unit  string `json:"unit"`
	Long string `json:"long"`
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
			"id":    id,
			"unit":  unit.Unit,
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		status := dbUnitsDELETE(id)
		c.Status(status)
	}
}