package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBoxes(t *testing.T) {
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
		{"GET auf /api/health", "GET", "/api/health", "", `{"version":1}`, 200},
		{"PUT auf /api/boxes", "PUT", "/api/boxes", `{"name": "Box1","notiz": "Notiz1"}`, `{"id":1,"name":"Box1","notiz":"Notiz1"}`, 201},
		{"GET auf /api/boxes", "GET", "/api/boxes", ``, `[{"id":1,"name":"Box1","notiz":"Notiz1"}]`, 200},
		{"PATCH auf /api/boxes", "PATCH", "/api/boxes/1", `{"name": "Patch1","notiz": "Patch1"}`, ``, 204},
		{"GET nach PATCH", "GET", "/api/boxes", ``, `[{"id":1,"name":"Patch1","notiz":"Patch1"}]`, 200},
		{"PUT auf /api/boxes", "PUT", "/api/boxes", `{"name": "box2","notiz": "notiz2"}`, `{"id":2,"name":"box2","notiz":"notiz2"}`, 201},
		{"DELETE auf /api/boxes", "DELETE", "/api/boxes/1", ``, ``, 204},
		{"GET nach DELETE", "GET", "/api/boxes", ``, `[{"id":2,"name":"box2","notiz":"notiz2"}]`, 200},
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