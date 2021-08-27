package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
	"log"
)

type Teacher struct {
	ID int `json:teacher_id`
	Surname string `json:surname`
	Name string `json:name`
	Patronymic string `json:patronymic`
	Post string `json:post`
	Education string `json:education`
	Qualification string `json:qualification`
}

type Student struct {
	ID int `json:student_id`
	Surname string `json:surname`
	Name string `json:name`
	Patronymic string `json:patronymic`
	DateBirth string `json:date_birth`
	ReceiptDate string `json:receipt_date`
	ExperationDate string `json:experation_date`
}

type Class struct {
	ID int `json:class_id`
	Name string `json:name`
	Student *Student `json:class_id`
	Teacher *Teacher `json:class_id`
}

var database *sql.DB

func Students(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("select * from simple_api.Student")
	if err != nil {
		log.Println(err);
	}
	defer rows.Close()
	students := []Student{}

	for rows.Next() {
		s := Student{}
		err := rows.Scan(&s.ID, &s.Surname, &s.Name, &s.Patronymic, &s.DateBirth, &s.ReceiptDate, &s.ExperationDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		students = append(students, s)
	}

	tmpl, _ := template.ParseFiles("templates/students.html")
	tmpl.Execute(w, students)
}

func Index(w http.ResponseWriter, r *http.Request) {
    data := "Index page"
    tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1><a href='/students'>Студенты</a>")
    tmpl.Execute(w, data)
}

func main() {
	db, err := sql.Open("mysql", "root:root@/simple_api")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	http.HandleFunc("/", Index)
	http.HandleFunc("/students", Students)

	log.Fatal(http.ListenAndServe(":8000", nil))
}