package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
	"log"
)

// Teacher struct
type Teacher struct {
	ID int `json:teacher_id`
	Surname string `json:surname`
	Name string `json:name`
	Patronymic string `json:patronymic`
	Post string `json:post`
	Education string `json:education`
	Qualification string `json:qualification`
}

// Student struct
type Student struct {
	ID int `json:student_id`
	Surname string `json:surname`
	Name string `json:name`
	Patronymic string `json:patronymic`
	DateBirth string `json:date_birth`
	ReceiptDate string `json:receipt_date`
	ExperationDate string `json:experation_date`
}

// Class struct
type Class struct {
	ID int `json:class_id`
	Name string `json:name`
	Student *Student `json:class_id`
	Teacher *Teacher `json:class_id`
}

var database *sql.DB

func Teachers(w http.ResponseWriter, r *http.Request) {
	// Get all teachers
	if r.Method == "GET" {
		// Get all students
		rows, err := database.Query("select * from simple_api.Student")
		if err != nil {
			log.Println(err);
		}
		defer rows.Close()
		// Array of students
		students := []Student{}

		// Append student in array
		for rows.Next() {
			s := Student{}
			err := rows.Scan(&s.ID, &s.Surname, &s.Name, &s.Patronymic, &s.DateBirth, &s.ReceiptDate, &s.ExperationDate)
			if err != nil {
				log.Println(err)
				continue
			}
			students = append(students, s)
		}

		// Output of templates
		tmpl, _ := template.ParseFiles("templates/students.html")
		tmpl.Execute(w, students)
	// Add teacher
	} else if r.Method == "POST" {
		// Parse form
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// Insert data in database
		_, err = database.Exec("insert into simple_api.Teacher (surname, name, patronymic, post, education, qualification) values (?,?,?,?,?,?)",
		r.FormValue("surname"), r.FormValue("name"), r.FormValue("patronymic"), r.FormValue("date_birth"), r.FormValue("receipt_date"), r.FormValue("expiration_date"))
		if err != nil {
			log.Println(err)
		}
		// Redirect
		http.Redirect(w, r, "/students", 301)
	}

}

func Students(w http.ResponseWriter, r *http.Request) {
	// Get all students
	if r.Method == "GET" {
		// Get all students
		rows, err := database.Query("select * from simple_api.Student")
		if err != nil {
			log.Println(err);
		}
		defer rows.Close()
		// Array of students
		students := []Student{}

		// Append student in array
		for rows.Next() {
			s := Student{}
			err := rows.Scan(&s.ID, &s.Surname, &s.Name, &s.Patronymic, &s.DateBirth, &s.ReceiptDate, &s.ExperationDate)
			if err != nil {
				log.Println(err)
				continue
			}
			students = append(students, s)
		}

		// Output of templates
		tmpl, _ := template.ParseFiles("templates/students.html")
		tmpl.Execute(w, students)
	// Add student
	} else if r.Method == "POST" {
		// Parse form
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// Insert data in database
		_, err = database.Exec("insert into simple_api.Student (surname, name, patronymic, date_birth, receipt_date, expiration_date) values (?,?,?,?,?,?)",
		r.FormValue("surname"), r.FormValue("name"), r.FormValue("patronymic"), r.FormValue("date_birth"), r.FormValue("receipt_date"), r.FormValue("expiration_date"))
		if err != nil {
			log.Println(err)
		}
		// Redirect
		http.Redirect(w, r, "/students", 301)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
    data := "Главная страница"
    tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1><a href='/students'>Студенты</a> | <a href='/teachers'>Преподаватели</a>")
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
	http.HandleFunc("/teachers", Teachers)

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}