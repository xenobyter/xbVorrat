package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Box struct {
	Name  string `json:"name"`
	Notiz string `json:"notiz"`
}

type Boxes []struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Notiz string `json:"notiz"`
}

func boxesPUT(c *gin.Context) {
	var box Box
	err := c.BindJSON(&box)
	if err != nil {
		log.Fatal(err)
	}

	if box.Name == "" {
		c.String(http.StatusBadRequest, "Name fehlt")
	} else {
		id := dbBoxesPUT(box)
		c.JSON(http.StatusCreated, gin.H{
			"id":    id,
			"name":  box.Name,
			"notiz": box.Notiz,
		})
	}
}

func boxesGET(c *gin.Context) {
	res := dbBoxesGET()
	c.JSON(http.StatusOK, res)
}

func boxesPATCH(c *gin.Context) {
	var box Box
	err := c.BindJSON(&box)
	if err != nil {
		log.Fatalf("boxesPATCH BindJSON: %v", err)
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		status := dbBoxesPATCH(id, box)
		c.Status(status)
	}
}

func boxesDELETE(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		status := dbDeleteByID("boxes", id)
		c.Status(status)
	}
}
