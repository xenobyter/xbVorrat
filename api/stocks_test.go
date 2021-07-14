package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_stocksPUT(t *testing.T) {
	setupDB()
	router := SetupRouter()

	dbBoxesPUT(Box{"Box", "Notiz"})
	dbUnitsPUT(Unit{"u", "Unit"})
	dbArticlesPUT(Article{"Article", 1})
	dbArticlesPUT(Article{"Article2", 1})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PUT", "PUT", "/api/stocks", `{"article":1,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2021"}`, `{"article":1,"box":1,"expiry":"31.12.2021","id":1,"quantity":1,"size":0.5}`, 201},
		{"PUT unbekannter Artikel", "PUT", "/api/stocks", `{"article":10,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2021"}`, `Unbekannter Artikel`, 409},
		{"PUT unbekannte Box", "PUT", "/api/stocks", `{"article":1,"box":10, "size":0.5,"quantity":1,"expiry":"31.12.2021"}`, `Unbekannte Box`, 409},
		{"PUT invalide Größe", "PUT", "/api/stocks", `{"article":1,"box":1, "size":"invalid","quantity":1,"expiry":"31.12.2021"}`, ``, 400},
		{"PUT invalide Anzahl", "PUT", "/api/stocks", `{"article":1,"box":1, "size":0.5,"quantity":"invalid","expiry":"31.12.2021"}`, ``, 400},
		{"PUT invalides Verfallsdatum", "PUT", "/api/stocks", `{"article":1,"box":1, "size":0.5,"quantity":1,"expiry":"invalid"}`, ``, 400},
		{"PUT zweiter Artikel", "PUT", "/api/stocks", `{"article":2,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2000"}`, `{"article":2,"box":1,"expiry":"31.12.2000","id":2,"quantity":1,"size":0.5}`, 201},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(tt.verb, tt.route, body)
			router.ServeHTTP(w, r)

			if tt.status != w.Code {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Code, tt.status)
			}
			if tt.res != w.Body.String() {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Body.String(), tt.res)
			}
		})
	}

	teardownDB()
}

func Test_stocksGET(t *testing.T) {
	setupDB()
	router := SetupRouter()

	dbBoxesPUT(Box{"Box", "Notiz"})
	dbUnitsPUT(Unit{"u", "Unit"})
	dbArticlesPUT(Article{"Article", 1})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"GET leere Liste", "GET", "/api/stocks", "", "null", 200},
		{"GET ein Artikel", "GET", "/api/stocks", "", `[{"id":1,"article":1,"box":1,"size":0.5,"quantity":2,"expiry":"31.12.2021"}]`, 200},
		{"GET zwei Artikel", "GET", "/api/stocks", "", `[{"id":1,"article":1,"box":1,"size":0.5,"quantity":2,"expiry":"31.12.2021"},{"id":2,"article":1,"box":1,"size":0.5,"quantity":2,"expiry":"31.12.2021"}]`, 200},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(tt.verb, tt.route, body)
			router.ServeHTTP(w, r)

			if tt.status != w.Code {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Code, tt.status)
			}
			if tt.res != w.Body.String() {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Body.String(), tt.res)
			}
			dbStocksPUT(Stock{1, 1, 0.5, 2, "31.12.2021"})
		})
	}
}

func Test_stocksPATCH(t *testing.T) {
	setupDB()
	router := SetupRouter()

	dbBoxesPUT(Box{"Box", "Notiz"})
	dbUnitsPUT(Unit{"u", "Unit"})
	dbArticlesPUT(Article{"Article", 1})
	dbArticlesPUT(Article{"Article2", 1})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PATCH leere Liste", "PATCH", "/api/stocks/1", `{"article":1,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2021"}`, "", 404},
		{"PATCH Verfallsdatum", "PATCH", "/api/stocks/1", `{"article":1,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2000"}`, "", 204},
		{"PATCH Anzahl Artikel 2", "PATCH", "/api/stocks/1", `{"article":1,"box":1, "size":0.5,"quantity":2,"expiry":"31.12.2021"}`, "", 204},
		{"PATCH Falsche Box", "PATCH", "/api/stocks/1", `{"article":1,"box":2, "size":0.5,"quantity":1,"expiry":"31.12.2000"}`, "Unbekannte Box", 409},
		{"PATCH Falscher Artikel", "PATCH", "/api/stocks/1", `{"article":3,"box":1, "size":0.5,"quantity":1,"expiry":"31.12.2000"}`, "Unbekannter Artikel", 409},
		{"PATCH invalides Verfallsdatum", "PATCH", "/api/stocks/1", `{"article":1,"box":1, "size":0.5,"quantity":1,"expiry":"invalid"}`, "", 400},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(tt.verb, tt.route, body)
			router.ServeHTTP(w, r)

			if tt.status != w.Code {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Code, tt.status)
			}
			if tt.res != w.Body.String() {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Body.String(), tt.res)
			}
			dbStocksPUT(Stock{1, 1, 0.5, 1, "31.12.2021"})
			dbStocksPUT(Stock{2, 1, 0.5, 1, "31.12.2021"})
		})
	}

	teardownDB()
}

func Test_expiryDateCheck(t *testing.T) {
	tests := []struct {
		name    string
		stock   Stock
		wantErr bool
	}{
		{"Einfaches Datum", Stock{Expiry: "01.01.2000"}, false},
		{"Striche", Stock{Expiry: "01-01-2000"}, true},
		{"Ohne Punkte", Stock{Expiry: "01012000"}, true},
		{"Jahr fehlt", Stock{Expiry: "01.01.00"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := expiryDateCheck(tt.stock); (err != nil) != tt.wantErr {
				t.Errorf("expiryDateCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_stocksDELETE(t *testing.T) {
	setupDB()
	router := SetupRouter()

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"DELETE leere Liste", "DELETE", "/api/stocks/1", ``, "", 404},
		{"DELETE ID 1", "DELETE", "/api/stocks/1", ``, "", 204},
		{"DELETE mit ungültiger ID", "DELETE", "/api/stocks/invalid", ``, "", 400},
		{"DELETE falscher ID", "DELETE", "/api/stocks/20", ``, "", 404},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(tt.verb, tt.route, body)
			router.ServeHTTP(w, r)

			if tt.status != w.Code {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Code, tt.status)
			}
			if tt.res != w.Body.String() {
				t.Errorf("%v auf %v ist: %v, soll %v", tt.verb, tt.route, w.Body.String(), tt.res)
			}
			dbStocksPUT(Stock{1, 1, 0.5, 1, "31.12.2021"})
		})
	}

	teardownDB()
}
