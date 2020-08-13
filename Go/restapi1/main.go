package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

type Shoe struct {
	Id int
	Shoe string
	ShoeDate string
	Description string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "db_user"
	dbPass := "jord@n_password"
	dbName := "shoes"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM ShoeNames ORDER BY id DESC ")
	if err != nil {
		panic(err.Error())
	}
	sh := Shoe{}
	res := []Shoe{}

	for selDB.Next() {
		var id int
		var shoe, shoedate, description string
		err = selDB.Scan(&id, &shoe, &shoedate, &description)
		if err != nil {
			panic(err.Error())
		}
		sh.Id = id
		sh.Shoe = shoe
		sh.ShoeDate = shoedate
		sh.Description = description
		res = append(res, sh)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM ShoeNames WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	sh := Shoe{}
	for selDB.Next() {
		var id int
		var shoe, shoedate, description string
		err = selDB.Scan(&id, &shoe, &shoedate, &description)
		if err != nil {
			panic(err.Error())
		}
		sh.Id = id
		sh.Shoe = shoe
		sh.ShoeDate = shoedate
		sh.Description = description
	}
	tmpl.ExecuteTemplate(w, "Show", sh)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	if r.Method == "POST"{
		shoe := r.FormValue("shoe")
		shoedate := r.FormValue("shoedate")
		description := r.FormValue("description")
		insForm, err := db.Prepare("INSERT INTO ShoeNames(shoe, shoedate,description) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(shoe, shoedate, description)
		log.Println("INSERT: Shoe: " + shoe + "| Shoe Date: "+ shoedate, " | Description: " + description)
	}
	defer db.Close()
	http.Redirect(w,r,"/", 301)
}

func Update(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	if r.Method == "POST" {
		shoe := r.FormValue("shoe")
		shoedate := r.FormValue("shoedate")
		description := r.FormValue("description")
		insForm, err := db.Prepare("UPDATE ShoeNames SET shoe=?, shoedate=?, description=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(shoe, shoedate, description)
		log.Println("UPDATE: Shoe: " + shoe + " | shoedate: " + shoedate + " | description "+ description)
	}
	defer db.Close()
	http.Redirect(w,r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	shoe := r.URL.Query().Get("id")
	defForm, err := db.Prepare("DELETE FROM ShoeNames where id=?")
	if err != nil {
		panic(err.Error())
	}
	defForm.Exec(shoe)
	log.Println("DELETE RECORD")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func getJSON(nId string)(string, error){
	db := dbConn()
	rows, err := db.Query("SELECT * FROM ShoeNames WHERE id=?", nId)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}

//func Edit(w http.ResponseWriter, r *http.Request){
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM ShoeNames where id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//
//}

func ViewJson(w http.ResponseWriter, r *http.Request){
	p:= r.URL.Query().Get("id")
	g, ok := getJSON(p)
	if ok != nil {
		fmt.Println("Problem")
	}
	fmt.Fprint(w, g)
}

func main() {
	log.Println("Jump Man started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/json", ViewJson)
	http.ListenAndServe(":8080", nil)
}


