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
		table string
		id    int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantBoxes  Boxes
	}{
		{"DELETE bei leerer DB", args{`boxes`, 1}, http.StatusNotFound, nil},
		{"DELETE Falsche id", args{"boxes", 10}, http.StatusNotFound, Boxes{{1, "name", "notiz"}}},
		{"DELETE ID 1", args{"boxes", 1}, http.StatusNoContent, Boxes{{2, "name", "notiz"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := dbDeleteByID(tt.args.table, tt.args.id); gotStatus != tt.wantStatus {
				t.Errorf("dbDeleteByID() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotBoxes := dbBoxesGET(); !reflect.DeepEqual(gotBoxes, tt.wantBoxes) {
				t.Errorf("dbDeleteByID() = %v, want %v", gotBoxes, tt.wantBoxes)
			}
		})
		dbBoxesPUT(Box{"name", "notiz"})
	}

	teardownDB()
}

func Test_dbUnitsPUT(t *testing.T) {
	setupDB()

	tests := []struct {
		name   string
		unit   Unit
		wantId int64
	}{
		{"Erste Einheit erstellen", Unit{"Box2", "Notiz2"}, 1},
		{"Zweite Einheit erstellen", Unit{"Box3", "Notiz3"}, 2},
	}
	for _, tt := range tests {
		log.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if gotId := dbUnitsPUT(tt.unit); gotId != tt.wantId {
				t.Errorf("dbUnitsPUT() = %v, want %v", gotId, tt.wantId)
			}
		})
	}

	teardownDB()
}

func TestUnitsGET(t *testing.T) {
	setupDB()

	tests := []struct {
		name string
		want Units
	}{
		{"Leere Liste", nil},
		{"Eine Einheit", Units{{1, "unit", "long"}}},
		{"Zweite Einheit", Units{{1, "unit", "long"}, {2, "unit", "long"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dbUnitsGET(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Units() = %v, want %v", got, tt.want)
			}
			dbUnitsPUT(Unit{"unit", "long"})
		})
	}

	teardownDB()
}

func Test_dbUnitsPATCH(t *testing.T) {
	setupDB()

	type args struct {
		id   int64
		unit Unit
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantUnit   string
		wantLong   string
	}{
		{"PATCH bei leerer DB", args{10, Unit{"neu", ""}}, http.StatusNotFound, "", ""},
		{"PATCH falsche ID", args{0, Unit{"neu", ""}}, http.StatusNotFound, "", ""},
		{"PATCH Einheit", args{1, Unit{"neu", ""}}, http.StatusNoContent, "neu", ""},
		{"PATCH Einheit und Langtext", args{1, Unit{"neu1", "neu1"}}, http.StatusNoContent, "neu1", "neu1"},
		{"PATCH Langtext leeren", args{1, Unit{"neu1", ""}}, http.StatusNoContent, "neu1", ""},
		{"PATCH Einheit nicht leeren", args{1, Unit{"", "neu"}}, http.StatusBadRequest, "neu1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := dbUnitsPATCH(tt.args.id, tt.args.unit)
			if status != tt.wantStatus {
				t.Errorf("dbUnitsPATCH() status = %v, wantStatus %v", status, tt.wantStatus)
			}
			units := dbUnitsGET()
			if len(units) != 0 && status != http.StatusNotFound {
				unit := units[tt.args.id-1]
				if unit.Unit != tt.wantUnit || unit.Long != tt.wantLong {
					t.Errorf("dbUnitsPATCH() unit = %v, long = %v, wantunit %v, WantLong %v", unit.Unit, unit.Long, tt.wantUnit, tt.wantLong)
				}
			} else {
				//db ist leer, einmalig eine box anlegen
				dbUnitsPUT(Unit{"unit", "long"})
			}
		})
	}

	teardownDB()
}

func Test_dbUnitsDELETE(t *testing.T) {
	setupDB()

	type args struct {
		table string
		id    int64
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantUnits  Units
	}{
		{"DELETE bei leerer DB", args{"units", 1}, http.StatusNotFound, nil},
		{"DELETE Falsche id", args{"units", 10}, http.StatusNotFound, Units{{1, "name", "notiz"}}},
		{"DELETE ID 1", args{"units", 1}, http.StatusNoContent, Units{{2, "name", "notiz"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := dbDeleteByID(tt.args.table, tt.args.id); gotStatus != tt.wantStatus {
				t.Errorf("dbUnitsDELETE() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotunits := dbUnitsGET(); !reflect.DeepEqual(gotunits, tt.wantUnits) {
				t.Errorf("dbUnitsDELETE() = %v, want %v", gotunits, tt.wantUnits)
			}
		})
		dbUnitsPUT(Unit{"name", "notiz"})
	}

	teardownDB()
}

func Test_dbArticlesPUT(t *testing.T) {
	setupDB()

	tests := []struct {
		name    string
		article Article
		wantId  int64
	}{
		{"Ersten Artikel anlegen", Article{"article", 1}, 1},
		{"Zweiten Artikel anlegen", Article{"article", 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := dbArticlesPUT(tt.article); gotId != tt.wantId {
				t.Errorf("dbArticlesPUT() = %v, want %v", gotId, tt.wantId)
			}
		})
	}

	teardownDB()
}

func Test_dbArticlesGET(t *testing.T) {
	setupDB()

	tests := []struct {
		name         string
		wantArticles Articles
	}{
		{"Leere Liste", nil},
		{"Ein Artikel", Articles{{1, "name", 1}}},
		{"Zwei Artikel", Articles{{1, "name", 1}, {2, "name", 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArticles := dbArticlesGET(); !reflect.DeepEqual(gotArticles, tt.wantArticles) {
				t.Errorf("dbArticlesGET() = %v, want %v", gotArticles, tt.wantArticles)
			}
		})
		dbArticlesPUT(Article{"name", 1})
	}

	teardownDB()
}

func Test_dbArticlesPATCH(t *testing.T) {
	setupDB()

	type args struct {
		id      int64
		article Article
	}
	tests := []struct {
		name         string
		args         args
		wantArticles Articles
		wantStatus   int
	}{
		{"PATCH bei leerer DB", args{0, Article{"name", 1}}, nil, http.StatusNotFound},
		{"PATCH korrekter Artikel", args{1, Article{"neu", 1}}, Articles{{1, "neu", 1}}, http.StatusNoContent},
		{"PATCH falsche ID", args{10, Article{"neu", 1}}, Articles{{1, "neu", 1}, {2, "name", 1}}, http.StatusNotFound},
		{"PATCH ohne Name", args{1, Article{"", 1}}, Articles{{1, "neu", 1}, {2, "name", 1}, {3, "name", 1}}, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := dbArticlesPATCH(tt.args.id, tt.args.article); gotStatus != tt.wantStatus {
				t.Errorf("dbArticlesPATCH() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotArticles := dbArticlesGET(); !reflect.DeepEqual(gotArticles, tt.wantArticles) {
				t.Errorf("dbArticlesPATCH() = %v, want %v", gotArticles, tt.wantArticles)
			}
		})
		dbArticlesPUT(Article{"name", 1})
	}

	teardownDB()
}

func Test_dbArticlesDELETE(t *testing.T) {
	setupDB()

	type args struct {
		table string
		id    int64
	}
	tests := []struct {
		name         string
		args         args
		wantStatus   int
		wantArticles Articles
	}{
		{"DELETE bei leerer DB", args{"articles", 1}, http.StatusNotFound, nil},
		{"DELETE korrekter Artikel", args{"articles", 1}, http.StatusNoContent, nil},
		{"DELETE falscher Artikel", args{"articles", 10}, http.StatusNotFound, Articles{{1, "name", 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus :=dbDeleteByID(tt.args.table,tt.args.id); gotStatus != tt.wantStatus {
				t.Errorf("dbArticlesDELETE() = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotArticles := dbArticlesGET(); !reflect.DeepEqual(gotArticles, tt.wantArticles) {
				t.Errorf("dbArticlesDELETE() = %v, want %v", gotArticles, tt.wantArticles)
			}
		})
		dbArticlesPUT(Article{"name", 1})
	}

	teardownDB()
}
