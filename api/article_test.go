package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_articlesPUT(t *testing.T) {
	setupDB()
	router := SetupRouter()
	unit := dbUnitsPUT(Unit{"u", "Unit"})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PUT", "PUT", "/api/articles", `{"name": "Name","unit": ` + fmt.Sprint(unit) + `}`, `{"id":1,"name":"Name","unit":` + fmt.Sprint(unit) + `}`, 201},
		{"PUT ohne Body", "PUT", "/api/articles", `{}`, `Artikel fehlt`, 400},
		{"PUT ohne Einheit", "PUT", "/api/articles", `{"name": "Name"}`, `Unbekannte Einheit`, 400},
		{"PUT falsche Einheit", "PUT", "/api/articles", `{"name": "Name","unit": 2}`, `Unbekannte Einheit`, 400},
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

func Test_articlesGET(t *testing.T) {
	setupDB()
	router := SetupRouter()
	dbUnitsPUT(Unit{"kg", "Kilogramm"})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"GET leere Liste", "GET", "/api/articles", "", "null", 200},
		{"GET ein Artikel", "GET", "/api/articles", "", `[{"id":1,"name":"name","unit":1}]`, 200},
		{"GET zwei Artikel", "GET", "/api/articles", "", `[{"id":1,"name":"name","unit":1},{"id":2,"name":"name","unit":1}]`, 200},
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
			dbArticlesPUT(Article{"name", 1})
		})
	}
}

func TestArticlesPatch(t *testing.T) {
	setupDB()
	router := SetupRouter()
	dbUnitsPUT(Unit{"kg", "Kilogramm"})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PATCH leere Liste", "PATCH", "/api/articles/1", `{"name": "Neu","unit": 1}`, "", 404},
		{"PATCH ok", "PATCH", "/api/articles/1", `{"name": "Neu","unit": 1}`, "", 204},
		{"PATCH ohne namen", "PATCH", "/api/articles/1", `{"unit": 1}`, "", 400},
		{"PATCH ohne Einheit", "PATCH", "/api/articles/1", `{"name": "Neu"}`, "", 409},
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
			dbArticlesPUT(Article{"name", 1})
		})
	}

	teardownDB()
}

func Test_articlesDELETE(t *testing.T) {
	setupDB()
	router := SetupRouter()
	dbUnitsPUT(Unit{"kg", "Kilogramm"})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"DELETE leere Liste", "DELETE", "/api/articles/1", ``, "", 404},
		{"DELETE ID 1", "DELETE", "/api/articles/1", ``, "", 204},
		{"DELETE mit ungültiger ID", "DELETE", "/api/articles/invalid", ``, "", 400},
		{"DELETE falscher ID", "DELETE", "/api/articles/20", ``, "", 404},
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
			dbArticlesPUT(Article{"name", 1})
		})
	}

	teardownDB()
}
