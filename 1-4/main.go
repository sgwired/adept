package main

import (
	"html/template"
	"net/http"
)

var port = "80"
var tpl *template.Template
var prod1 product

var name = "John Doe"

type User struct {
	Name     string
	Language string
	Member   bool
}

type product struct {
	ID    int
	Name  string
	Cost  float32
	Specs prodSpec
}

type prodSpec struct {
	Size   string
	Weight float32
	Desc   string
}

var u User

func main() {

	// // tpl = template.Must(template.ParseFiles("templates/index2.html"))
	// // tpl, _ = template.ParseFiles("templates/index2.html")
	// tpl, _ = tpl.ParseFiles("templates/index2.html")

	// u = User{"John Doe", "English", false}
	// u = User{"Jose Doe", "Spanish", true}
	u = User{"Ja Doe", "Mandarin", true}

	prod1 = product{
		ID:   1,
		Name: "Apple",
		Cost: 100.00,
		Specs: prodSpec{
			Size:   "10",
			Weight: 200.00,
			Desc:   "This is an apple",
		},
	}

	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	http.HandleFunc("/contact", contactHandleFunc)
	http.HandleFunc("/login", loginHandleFunc)
	http.HandleFunc("/welcome", welcomeHander)

	http.ListenAndServe(":"+port, nil)
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "r.URL.Path = %q\n", r.URL.Path)

	tpl.ExecuteTemplate(w, "index3.html", prod1)
}

func welcomeHander(w http.ResponseWriter, r *http.Request) {

	// tpl.ExecuteTemplate(w, "membership.html", u)
	tpl.ExecuteTemplate(w, "membership2.html", u)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandleFunc(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func loginHandleFunc(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "login.html", nil)
}
