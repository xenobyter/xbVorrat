package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var (
	db *sql.DB
)

func createTables() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS boxes (name TEXT, notiz TEXT);")
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS units (unit TEXT, long TEXT);")
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS articles (name TEXT, unit INTEGER);")
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS stocks (article INTEGER, box INTEGER, size REAL, quantity INTEGER, expiry TEXT );")
	if err != nil {
		log.Panic(err)
	}
}

func dbBoxesPUT(box Box) (id int64) {
	result, err := db.Exec("INSERT INTO boxes (name, notiz) VALUES (?, ?);", box.Name, box.Notiz)
	if err != nil {
		log.Fatalf("Error in INSERT INTO boxes: %v", err)
	}
	id, _ = result.LastInsertId()
	return
}

func dbBoxesGET() (boxes Boxes) {
	box := make(Boxes, 1)
	queryStmt := "SELECT rowid, name, notiz FROM boxes;"
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&box[0].ID, &box[0].Name, &box[0].Notiz)
		if err != nil {
			log.Fatalf("Error in Scanning Rows:")
		}
		boxes = append(boxes, box[0])
	}
	return boxes
}

func dbBoxesPATCH(id int64, box Box) int {
	if box.Name != "" {
		res, err := db.Exec("UPDATE boxes SET name = ?, notiz = ? WHERE rowid = ?", box.Name, box.Notiz, id)
		cnt, _ := res.RowsAffected()
		switch {
		case err != nil:
			return http.StatusBadRequest
		case cnt == 0:
			return http.StatusNotFound
		case cnt == 1:
			return http.StatusNoContent
		}
	}
	return http.StatusBadRequest
}

func dbUnitsPUT(unit Unit) (id int64) {
	result, err := db.Exec("INSERT INTO units (unit, long) VALUES (?, ?);", unit.Unit, unit.Long)
	if err != nil {
		log.Fatalf("Error in INSERT INTO units: %v", err)
	}
	id, _ = result.LastInsertId()
	return
}

func dbUnitsGET() (units Units) {
	box := make(Units, 1)
	queryStmt := "SELECT rowid, unit, long FROM units;"
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&box[0].ID, &box[0].Unit, &box[0].Long)
		if err != nil {
			log.Fatalf("Error in Scanning Rows:")
		}
		units = append(units, box[0])
	}
	return units
}

func dbUnitsPATCH(id int64, unit Unit) int {
	if unit.Unit != "" {
		res, err := db.Exec("UPDATE units SET unit = ?, long = ? WHERE rowid = ?", unit.Unit, unit.Long, id)
		cnt, _ := res.RowsAffected()
		switch {
		case err != nil:
			return http.StatusBadRequest
		case cnt == 0:
			return http.StatusNotFound
		case cnt == 1:
			return http.StatusNoContent
		}
	}
	return http.StatusBadRequest
}

func dbArticlesPUT(article Article) (id int64) {
	result, err := db.Exec("INSERT INTO articles (name, unit) VALUES (?, ?);", article.Name, article.UnitID)
	if err != nil {
		log.Fatalf("Error in INSERT INTO articles: %v", err)
	}
	id, _ = result.LastInsertId()
	return
}

func dbArticlesGET() (articles Articles) {
	article := make(Articles, 1)
	queryStmt := "SELECT articles.rowid, name, unit, COALESCE(quantity * size,0) FROM articles LEFT JOIN stocks on articles.rowid = stocks.article;"
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&article[0].ID, &article[0].Name, &article[0].UnitID, &article[0].Quantity)
		if err != nil {
			log.Fatalf("Error in Scanning Rows:")
		}
		articles = append(articles, article[0])
	}
	return articles
}

func dbArticlesPATCH(id int64, article Article) int {
	if article.Name != "" {
		res, err := db.Exec("UPDATE articles SET name = ?, unit = ? WHERE rowid = ?", article.Name, article.UnitID, id)
		cnt, _ := res.RowsAffected()
		switch {
		case err != nil:
			return http.StatusBadRequest
		case cnt == 0:
			return http.StatusNotFound
		case cnt == 1:
			return http.StatusNoContent
		}
	}
	return http.StatusBadRequest
}

func dbDeleteByID(table string, id int64) int {
	sql := fmt.Sprint("DELETE FROM ", table, " WHERE rowid = ?")
	res, err := db.Exec(sql, id)
	if err != nil {
		log.Fatal(err)
		return http.StatusBadRequest
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return http.StatusNotFound
	}
	return http.StatusNoContent

}

func dbStocksPUT(stock Stock) (id int64) {
	result, err := db.Exec("INSERT INTO stocks (article, box, size, quantity, expiry) VALUES (?, ?, ?, ?, ?);", stock.Article, stock.Box, stock.Size, stock.Quantity, stock.Expiry)
	if err != nil {
		log.Fatalf("Error in INSERT INTO stocks: %v", err)
	}
	id, _ = result.LastInsertId()
	return
}

func dbStocksGET() (stocks Stocks) {
	stock := make(Stocks, 1)
	queryStmt := "SELECT rowid, article, box, size, quantity, expiry FROM stocks;"
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&stock[0].ID, &stock[0].Article, &stock[0].Box, &stock[0].Size, &stock[0].Quantity, &stock[0].Expiry)
		if err != nil {
			log.Fatalf("Error in Scanning Rows:")
		}
		stocks = append(stocks, stock[0])
	}
	return stocks
}
func dbStocksRichGET(aSort, aOrder string) (stocks StocksRich) {
	mSort := map[string]string{"id": "stocks.rowid", "articlestr": "articles.name", "boxstr": "boxes.name", "expiry": "SubStr(expiry,7,4)||SubStr(expiry,4,2)||SubStr(expiry,1,2)"}
	mOrder := map[string]string{"asc": "ASC", "desc": "DESC"}
	qSort, ok := mSort[aSort]
	if !ok {
		qSort = mSort["id"]
	}
	qOrder, ok := mOrder[aOrder]
	if !ok {
		qOrder = mOrder["asc"]
	}

	stock := make(StocksRich, 1)
	queryStmt := fmt.Sprintf("SELECT stocks.rowid, article, articles.name, box, boxes.name, size, units.unit, quantity, expiry from stocks INNER JOIN articles on articles.rowid = stocks.article INNER JOIN boxes on boxes.rowid = stocks.box INNER JOIN units on units.rowid = articles.unit ORDER BY %v %v;", qSort, qOrder)
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}

	for rows.Next() {
		err = rows.Scan(&stock[0].ID, &stock[0].Article, &stock[0].ArticleName, &stock[0].Box, &stock[0].BoxName, &stock[0].Size, &stock[0].Unit, &stock[0].Quantity, &stock[0].Expiry)
		if err != nil {
			log.Fatalf("Error in Scanning Rows:")
		}
		stocks = append(stocks, stock[0])
	}
	return stocks
}

func dbStocksPATCH(id int64, stock Stock) int {
	if stock.Expiry != "" {
		res, err := db.Exec("UPDATE stocks SET article = ?, box = ?, size = ?, quantity = ?, expiry = ? WHERE rowid = ?", stock.Article, stock.Box, stock.Size, stock.Quantity, stock.Expiry, id)
		cnt, _ := res.RowsAffected()
		switch {
		case err != nil:
			return http.StatusBadRequest
		case cnt == 0:
			return http.StatusNotFound
		case cnt == 1:
			return http.StatusNoContent
		}
	}
	return http.StatusBadRequest
}
