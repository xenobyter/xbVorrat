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

func (b Boxes) contains(id int64) bool {
	for _, n := range b {
		if id == n.ID {
			return true
		}
	}
	return false
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
	stocks := dbStocksGET()
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	switch {
	case parseErr != nil:
		c.Status(http.StatusBadRequest)
	case stocks.containsBox(id):
		c.String(http.StatusForbidden, "Box muss zum löschen leer sein")
	default:
		status := dbDeleteByID("boxes", id)
		c.Status(status)
	}
}
