package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Phrase struct {
	ID          int
	Jack        int
	Native      string
	Translation string
	CourseName  string
	Category    string
	NatLanguage string
	Counter     int
	Correct     int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "flashcardsdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM phrase ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	phr := Phrase{}
	res := []Phrase{}
	for selDB.Next() {
		var id, jack, counter, correct int
		var native, translation, courseName, category, natLanguage string
		err = selDB.Scan(&id, &jack, &native, &translation, &courseName, &category, &natLanguage, &counter, &correct)
		if err != nil {
			panic(err.Error())
		}
		phr.ID = id
		phr.Jack = jack
		phr.Native = native
		phr.Translation = translation
		phr.CourseName = courseName
		phr.Category = category
		phr.NatLanguage = natLanguage
		phr.Counter = counter
		phr.Correct = correct
		res = append(res, phr)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func ListCourses(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT DISTINCT courseName From phrase;")
	if err != nil {
		panic(err.Error())
	}
	phr := Phrase{}
	res := []Phrase{}
	for selDB.Next() {
		var courseName string
		err = selDB.Scan(&courseName)
		if err != nil {
			panic(err.Error())
		}
		phr.CourseName = courseName
		res = append(res, phr)
	}
	tmpl.ExecuteTemplate(w, "Courses", res)
	defer db.Close()
}

func Learn(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("courseName")
	selDB, err := db.Query("SELECT native, translation, correct, counter FROM phrase WHERE courseName=? ORDER BY RAND() LIMIT 1", nID)
	if err != nil {
		panic(err.Error())
	}
	phr := Phrase{}
	for selDB.Next() {
		var counter, correct int
		var native, translation string
		err = selDB.Scan(&native, &translation, &correct, &counter)
		if err != nil {
			panic(err.Error())
		}
		phr.Native = native
		phr.Translation = translation
		phr.Correct = correct
		phr.Counter = counter
		phr.CourseName = nID
	}
	tmpl.ExecuteTemplate(w, "Learn", phr)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phrase WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	phr := Phrase{}
	for selDB.Next() {
		var id, jack, counter, correct int
		var native, translation, courseName, category, natLanguage string
		err = selDB.Scan(&id, &jack, &native, &translation, &courseName, &category, &natLanguage, &counter, &correct)
		if err != nil {
			panic(err.Error())
		}
		phr.ID = id
		phr.Jack = jack
		phr.Native = native
		phr.Translation = translation
		phr.CourseName = courseName
		phr.Category = category
		phr.NatLanguage = natLanguage
		phr.Counter = counter
		phr.Correct = correct
	}
	tmpl.ExecuteTemplate(w, "Show", phr)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phrase WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	phr := Phrase{}
	for selDB.Next() {
		var id, jack, counter, correct int
		var native, translation, courseName, category, natLanguage string
		err = selDB.Scan(&id, &jack, &native, &translation, &courseName, &category, &natLanguage, &counter, &correct)
		if err != nil {
			panic(err.Error())
		}
		phr.ID = id
		phr.Jack = jack
		phr.Native = native
		phr.Translation = translation
		phr.CourseName = courseName
		phr.Category = category
		phr.NatLanguage = natLanguage
		phr.Counter = counter
		phr.Correct = correct
	}
	tmpl.ExecuteTemplate(w, "Edit", phr)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		jack := r.FormValue("jack")
		native := r.FormValue("native")
		translation := r.FormValue("translation")
		courseName := r.FormValue("courseName")
		category := r.FormValue("category")
		natLanguage := r.FormValue("natLanguage")
		insForm, err := db.Prepare("INSERT INTO phrase(jack, native, translation, courseName, category, natLanguage) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(jack, native, translation, courseName, category, natLanguage)
		log.Println("INSERT: Jack: " + jack + " | Native: " + native + " | Translation: " + translation + " | courseName: " + courseName + " | Category: " + category + " | NatLanguage: " + natLanguage)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		jack := r.FormValue("jack")
		native := r.FormValue("native")
		translation := r.FormValue("translation")
		courseName := r.FormValue("courseName")
		category := r.FormValue("category")
		natLanguage := r.FormValue("natLanguage")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE phrase SET jack=?, native=?, translation=?, courseName=?, category=?, natLanguage=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(jack, native, translation, courseName, category, natLanguage, id)
		log.Println("UPDATE: Jack: " + jack + " | Native: " + native + " | Translation: " + translation + " | courseName: " + courseName + " | Category: " + category + " | NatLanguage: " + natLanguage)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	phr := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM phrase WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(phr)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/courses", ListCourses)
	http.HandleFunc("/learn", Learn)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
