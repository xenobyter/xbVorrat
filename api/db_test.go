package api

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// setup

var tmp string

func setupDB() {
	var err error
	tmp, _ = ioutil.TempDir("", "xbVorrat")
	db, err = sql.Open("sqlite3", tmp+"/.xbVorrat")
	if err != nil {
		log.Panic(err)
	}
	createTables()
}

func teardownDB() {
	os.RemoveAll(tmp)
}

func Test_boxesPUT(t *testing.T) {
	setupDB()

	tests := []struct {
		name   string
		box    Box
		wantId int64
	}{
		{"Erste Box erstellen", Box{"Box2", "Notiz2"}, 1},
		{"Zweite Box erstellen", Box{"Box3", "Notiz3"}, 2},
	}
	for _, tt := range tests {
		log.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if gotId := dbBoxesPUT(tt.box); gotId != tt.wantId {
				t.Errorf("boxesPUT() = %v, want %v", gotId, tt.wantId)
			}
		})
	}

	teardownDB()
}

func TestBoxesGET(t *testing.T) {
	setupDB()

	tests := []struct {
		name string
		want Boxes
	}{
		{"Leere Liste", nil},
		{"Eine Box", Boxes{{1, "name", "notiz"}}},
		{"Zweite Box", Boxes{{1, "name", "notiz"}, {2, "name", "notiz"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dbBoxesGET(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoxesGET() = %v, want %v", got, tt.want)
			}
			dbBoxesPUT(Box{"name", "notiz"})
		})
	}

	teardownDB()
}

func Test_dbBoxesPATCH(t *testing.T) {
	setupDB()

	type args struct {
		id  int64
		box Box
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantName   string
		wantNotiz  string
	}{
		{"PATCH bei leerer DB", args{10, Box{"neu", ""}}, http.StatusNotFound, "", ""},
		{"PATCH falsche ID", args{0, Box{"neu", ""}}, http.StatusNotFound, "", ""},
		{"PATCH Name", args{1, Box{"neu", ""}}, http.StatusNoContent, "neu", ""},
		{"PATCH Name und Notiz", args{1, Box{"neu1", "neu1"}}, http.StatusNoContent, "neu1", "neu1"},
		{"PATCH Notiz leeren", args{1, Box{"neu1", ""}}, http.StatusNoContent, "neu1", ""},
		{"PATCH Name nicht leeren", args{1, Box{"", "neu"}}, http.StatusBadRequest, "neu1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := dbBoxesPATCH(tt.args.id, tt.args.box)
			if status != tt.wantStatus {
				t.Errorf("dbBoxesPATCH() status = %v, wantStatus %v", status, tt.wantStatus)
			}
			boxes := dbBoxesGET()
			if len(boxes) != 0 && status != http.StatusNotFound {
				box := boxes[tt.args.id-1]
				if box.Name != tt.wantName || box.Notiz != tt.wantNotiz {
					t.Errorf("dbBoxesPATCH() name = %v, notiz = %v, wantName %v, WantNotiz %v", box.Name, box.Notiz, tt.wantName, tt.wantNotiz)
				}
			} else {
				//db ist leer, einmalig eine box anlegen
				dbBoxesPUT(Box{"name", "notiz"})
			}
		})
	}

	teardownDB()
}

func Test_dbBoxesDELETE(t *testing.T) {
	setupDB()

	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantBoxes  Boxes
	}{
		{"DELETE bei leerer DB", args{1}, http.StatusNotFound, nil},
		{"DELETE Falsche id", args{10}, http.StatusNotFound, Boxes{{1, "name", "notiz"}}},
		{"DELETE ID 1", args{1}, http.StatusNoContent, Boxes{{2, "name", "notiz"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := dbBoxesDELETE(tt.args.id); gotStatus != tt.wantStatus {
				t.Errorf("dbBoxesDELETE() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotBoxes := dbBoxesGET(); !reflect.DeepEqual(gotBoxes, tt.wantBoxes) {
				t.Errorf("dbBoxesDELETE() = %v, want %v", gotBoxes, tt.wantBoxes)
			}
		})
		dbBoxesPUT(Box{"name", "notiz"})
	}

	teardownDB()
}
