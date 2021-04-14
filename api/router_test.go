package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name   string
		verb   string
		route  string
		body   io.Reader
		res    string
		status int
	}{
		{"GET auf /api/health", "GET", "/api/health", nil, `{"version":1}`, 200},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(tt.verb, tt.route, tt.body)
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
