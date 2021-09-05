package main

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
	"encoding/json"
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

// Get all teachers
func GetTeachers() []Teacher {
	// Get all teachers
	rows, err := database.Query("select * from simple_api.Teacher")
	if err != nil {
		log.Println(err);
	}
	defer rows.Close()
	// Array of teachers
	teachers := []Teacher{}

	// Append teacher in array
	for rows.Next() {
		t := Teacher{}
		err := rows.Scan(&t.ID, &t.Surname, &t.Name, &t.Patronymic, &t.Post, &t.Education, &t.Qualification)
		if err != nil {
			log.Println(err)
			continue
		}
		teachers = append(teachers, t)
	}
	return teachers
}

// Get all students
func GetStudents() []Student {
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
	return students
}

func Teachers(w http.ResponseWriter, r *http.Request) {
	// Get all teachers
	if r.Method == "GET" {
		// Get all teachers
		teachers := GetTeachers()
		// Output of templates
		tmpl, _ := template.ParseFiles("public/templates/teachers.html")
		tmpl.Execute(w, teachers)
	// Add teacher
	} else if r.Method == "POST" {
		// Parse form
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// Insert data in database
		_, err = database.Exec("insert into simple_api.Teacher (surname, name, patronymic, post, education, qualification) values (?,?,?,?,?,?)",
		r.FormValue("surname"), r.FormValue("name"), r.FormValue("patronymic"), r.FormValue("post"), r.FormValue("education"), r.FormValue("qualification"))
		if err != nil {
			log.Println(err)
		}
		// Redirect
		http.Redirect(w, r, "/teachers", 301)
	}

}

func Students(w http.ResponseWriter, r *http.Request) {
	// Get all students
	if r.Method == "GET" {
		// Get all students
		students := GetStudents()
		// Output of templates
		tmpl, _ := template.ParseFiles("public/templates/students.html")
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

// Main page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("public/index.html")
	tmpl.Execute(w, "index");
}

// Api teachers
func ApiTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		teachers := GetTeachers()
		json.NewEncoder(w).Encode(teachers)
	} else if r.Method == "POST" {

	}
}

// Api students
func ApiStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		students := GetStudents()
		json.NewEncoder(w).Encode(students)
	} else if r.Method == "POST" {
		
	}
}

// Api classes
func ApiClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	// Connection to base
	db, err := sql.Open("mysql", "root:root@/simple_api")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	// For connecting files
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Routes
	http.HandleFunc("/", Index)
	http.HandleFunc("/students", Students)
	http.HandleFunc("/teachers", Teachers)

	// Api routes
	http.HandleFunc("/api/teachers", ApiTeachers)
	http.HandleFunc("/api/students", ApiStudents)
	http.HandleFunc("/api/classes", ApiClasses)

	// Start local server
	log.Println("Запуск веб-сервера на http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}