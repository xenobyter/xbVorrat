package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestPUTBoxes(t *testing.T) {
	setupDB()
	defer teardownDB()
	router := SetupRouter()

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PUT", "PUT", "/api/boxes", `{"name": "Box1","notiz": "Notiz1"}`, `{"id":1,"name":"Box1","notiz":"Notiz1"}`, 201},
		{"PUT ohne Name", "PUT", "/api/boxes", `{"notiz": "Notiz1"}`, `Name fehlt`, 400},
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

}

func TestGETBoxes(t *testing.T) {
	setupDB()
	defer teardownDB()
	router := SetupRouter()

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"GET leere Liste", "GET", "/api/boxes", "", "null", 200},
		{"GET eine Box", "GET", "/api/boxes", "", `[{"id":1,"name":"name","notiz":"notiz"}]`, 200},
		{"GET zwei Boxen", "GET", "/api/boxes", "", `[{"id":1,"name":"name","notiz":"notiz"},{"id":2,"name":"name","notiz":"notiz"}]`, 200},
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
			dbBoxesPUT(Box{"name", "notiz"})
		})
	}

}

func TestPATCHBoxes(t *testing.T) {
	setupDB()
	defer teardownDB()
	router := SetupRouter()

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"PATCH leere Liste", "PATCH", "/api/boxes/1", `{"name": "Box1","notiz": "Notiz1"}`, "", 404},
		{"PATCH ID 1", "PATCH", "/api/boxes/1", `{"name": "Patch1","notiz": "Patch1"}`, "", 204},
		{"PATCH ID ohne Name", "PATCH", "/api/boxes/1", `{"name": "","notiz": "Patch1"}`, "", 400},
		{"PATCH mit ungültiger ID", "PATCH", "/api/boxes/invalid", `{"name": "Patch1","notiz": "Patch1"}`, "", 400},
		{"PATCH falscher ID", "PATCH", "/api/boxes/20", `{"name": "Patch1","notiz": "Patch1"}`, "", 404},
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
			dbBoxesPUT(Box{"name", "notiz"})
		})
	}

}

func TestDELETEBoxes(t *testing.T) {
	setupDB()
	defer teardownDB()
	router := SetupRouter()
	dbStocksPUT(Stock{1,2,1,1,"31.12.2021"})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"DELETE leere Liste", "DELETE", "/api/boxes/1", ``, "", 404},
		{"DELETE ID 1", "DELETE", "/api/boxes/1", ``, "", 204},
		{"DELETE mit ungültiger ID", "DELETE", "/api/boxes/invalid", ``, "", 400},
		{"DELETE falscher ID", "DELETE", "/api/boxes/20", ``, "", 404},
		{"DELETE nicht leer Box", "DELETE", "/api/boxes/2", ``, "Box muss zum löschen leer sein", 403},
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
			dbBoxesPUT(Box{"name", "notiz"})
		})
	}

}