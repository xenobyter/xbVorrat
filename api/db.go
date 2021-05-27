package api

import (
	"database/sql"
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

func dbBoxesDELETE(id int64) int {
	res, err := db.Exec("DELETE FROM boxes WHERE rowid = ?", id)
	cnt, _ := res.RowsAffected()
	switch {
	case err != nil:
		return http.StatusBadRequest
	case cnt == 0:
		return http.StatusNotFound
	case cnt == 1:
		return http.StatusNoContent
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

func dbUnitsDELETE(id int64) int {
	res, err := db.Exec("DELETE FROM units WHERE rowid = ?", id)
	cnt, _ := res.RowsAffected()
	switch {
	case err != nil:
		return http.StatusBadRequest
	case cnt == 0:
		return http.StatusNotFound
	case cnt == 1:
		return http.StatusNoContent
	}
	return http.StatusBadRequest
}