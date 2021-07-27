package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPUTUnits(t *testing.T) {
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
		{"PUT", "PUT", "/api/units", `{"unit": "Unit1","long": "Long1"}`, `{"id":1,"long":"Long1","unit":"Unit1"}`, 201},
		{"PUT ohne Unit", "PUT", "/api/units", `{"long": "Long"}`, `Unit fehlt`, 400},
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

func TestGETUnits(t *testing.T) {
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
		{"GET leere Liste", "GET", "/api/units", "", "null", 200},
		{"GET eine Box", "GET", "/api/units", "", `[{"id":1,"unit":"unit","long":"long"}]`, 200},
		{"GET zwei Boxen", "GET", "/api/units", "", `[{"id":1,"unit":"unit","long":"long"},{"id":2,"unit":"unit","long":"long"}]`, 200},
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
			dbUnitsPUT(Unit{"unit", "long"})
		})
	}
}

func TestPATCHUnits(t *testing.T) {
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
		{"PATCH leere Liste", "PATCH", "/api/units/1", `{"unit": "Box1","long": "Notiz1"}`, "", 404},
		{"PATCH ID 1", "PATCH", "/api/units/1", `{"unit": "Patch1","long": "Patch1"}`, "", 204},
		{"PATCH ID ohne Einheit", "PATCH", "/api/units/1", `{"unit": "","long": "Patch1"}`, "", 400},
		{"PATCH mit ungültiger ID", "PATCH", "/api/units/invalid", `{"unit": "Patch1","long": "Patch1"}`, "", 400},
		{"PATCH falscher ID", "PATCH", "/api/units/20", `{"unit": "Patch1","long": "Patch1"}`, "", 404},
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
			dbUnitsPUT(Unit{"name", "notiz"})
		})
	}

}

func TestDELETEUnits(t *testing.T) {
	setupDB()
	defer teardownDB()
	router := SetupRouter()
	dbArticlesPUT(Article{"name", 2})

	tests := []struct {
		name   string
		verb   string
		route  string
		body   string
		res    string
		status int
	}{
		{"DELETE leere Liste", "DELETE", "/api/units/1", ``, "", 404},
		{"DELETE ID 1", "DELETE", "/api/units/1", ``, "", 204},
		{"DELETE mit ungültiger ID", "DELETE", "/api/units/invalid", ``, "", 400},
		{"DELETE falscher ID", "DELETE", "/api/units/20", ``, "", 404},
		{"DELETE Einheit wird in Artikel benutzt", "DELETE", "/api/units/2", ``, "", 409},
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
			dbUnitsPUT(Unit{"name", "notiz"})
		})
	}

}

func TestUnits_contains(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		u    Units
		args args
		want bool
	}{
		{"Einfaches true", Units{{1, "kg", "Kilogramm"}}, args{1}, true},
		{"Einfaches false", Units{{1, "kg", "Kilogramm"}}, args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.contains(tt.args.id); got != tt.want {
				t.Errorf("Units.contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
