package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Name   string `json:"name"`
	UnitID int64  `json:"unit"`
}

type Articles []struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UnitID int64  `json:"unit"`
}

func (a Articles) contains(id int64) bool {
	for _, n := range a {
		if id == n.ID {
			return true
		}
	}
	return false
}

func (a Articles) containsUnit(id int64) bool {
	for _, n := range a {
		if id == n.UnitID {
			return true
		}
	}
	return false
}

func articlesPUT(c *gin.Context) {
	var article Article
	err := c.BindJSON(&article)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	units := dbUnitsGET()

	switch {
	case article.Name == "":
		c.String(http.StatusBadRequest, "Artikel fehlt")
	case !units.contains(article.UnitID):
		c.String(http.StatusBadRequest, "Unbekannte Einheit")
	default:
		id := dbArticlesPUT(article)
		c.JSON(http.StatusCreated, gin.H{
			"id":   id,
			"name": article.Name,
			"unit": article.UnitID,
		})
	}
}

func articlesGET(c *gin.Context) {
	res := dbArticlesGET()
	c.JSON(http.StatusOK, res)
}

func articlesPATCH(c *gin.Context) {
	var article Article
	err := c.BindJSON(&article)
	if err != nil {
		c.Status((http.StatusBadRequest))
	}
	units := dbUnitsGET()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	switch {
	case err != nil:
		c.Status(http.StatusBadRequest)
	case article.Name == "":
		c.Status(http.StatusBadRequest)
	case !units.contains(article.UnitID):
		c.Status(http.StatusConflict)
	default:
		status := dbArticlesPATCH(id, article)
		c.Status(status)
	}
}

func articlesDELETE(c *gin.Context) { 
	stocks := dbStocksGET()
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	switch {
	case parseErr != nil:
		c.Status(http.StatusBadRequest)
		case stocks.containsArticle(id):
		c.String(http.StatusForbidden, "Artikel noch in Verwendung")
	default:
		status := dbDeleteByID("articles", id)
		c.Status(status)
	}
}
