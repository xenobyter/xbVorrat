package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Stock struct {
	Article  int64   `json:"article"`
	Box      int64   `json:"box"`
	Size     float64 `json:"size"`
	Quantity int64   `json:"quantity"`
	Expiry   string  `json:"expiry"`
}

type Stocks []struct {
	ID       int64   `json:"id"`
	Article  int64   `json:"article"`
	Box      int64   `json:"box"`
	Size     float64 `json:"size"`
	Quantity int64   `json:"quantity"`
	Expiry   string  `json:"expiry"`
}

type StocksRich []struct {
	ID          int64   `json:"id"`
	Article     int64   `json:"article"`
	Box         int64   `json:"box"`
	Size        float64 `json:"size"`
	Quantity    int64   `json:"quantity"`
	Expiry      string  `json:"expiry"`
	ArticleName string  `json:"articlestr"`
	BoxName     string  `json:"boxstr"`
	Unit        string  `json:"unitstr"`
	Expired     bool    `json:"expired"`
}

func (s Stocks) containsBox(id int64) bool {
	for _, n := range s {
		if id == n.Box {
			return true
		}
	}
	return false
}
func (s Stocks) containsArticle(id int64) bool {
	for _, n := range s {
		if id == n.Article {
			return true
		}
	}
	return false
}

func stocksPUT(c *gin.Context) {
	articles := dbArticlesGET()
	boxes := dbBoxesGET()

	var stock Stock
	bindErr := c.BindJSON(&stock)

	switch {
	case bindErr != nil:
		c.Status(http.StatusBadRequest)
	case stock.Expiry == "":
		c.String(http.StatusBadRequest, "Verfallsdatum fehlt")
	case !articles.contains(stock.Article):
		c.String(http.StatusConflict, "Unbekannter Artikel")
	case !boxes.contains(stock.Box):
		c.String(http.StatusConflict, "Unbekannte Box")
	case expiryDateCheck(stock) != nil:
		c.Status(http.StatusBadRequest)
	default:
		id := dbStocksPUT(stock)
		c.JSON(http.StatusCreated, gin.H{
			"id":       id,
			"article":  stock.Article,
			"box":      stock.Box,
			"size":     stock.Size,
			"quantity": stock.Quantity,
			"expiry":   stock.Expiry,
		})
	}
}

func expiryDateCheck(stock Stock) error {
	layout := "02.01.2006"
	_, datErr := time.Parse(layout, stock.Expiry)
	return datErr
}

func stocksRichGET(c *gin.Context) {
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")
	res := dbStocksRichGET(sort, order)
	c.JSON(http.StatusOK, res)
}

func stocksPATCH(c *gin.Context) {
	var stock Stock
	bindErr := c.BindJSON(&stock)
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)

	boxes := dbBoxesGET()
	articles := dbArticlesGET()

	switch {
	case bindErr != nil || parseErr != nil:
		c.Status(http.StatusBadRequest)
	case !boxes.contains(stock.Box):
		c.String(http.StatusConflict, "Unbekannte Box")
	case !articles.contains(stock.Article):
		c.String(http.StatusConflict, "Unbekannter Artikel")
	case expiryDateCheck(stock) != nil:
		c.Status(http.StatusBadRequest)
	default:
		status := dbStocksPATCH(id, stock)
		c.Status(status)
	}
}

func stocksDELETE(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		status := dbDeleteByID("stocks", id)
		c.Status(status)
	}
}
