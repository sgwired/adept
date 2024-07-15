package main

import (
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}
type GroceryList []string

var tpl *template.Template

var g GroceryList

var Todo []task

func main() {
	g = GroceryList{"milk", "bread", "eggs", "cheese", "yogurt", "sugar"}
	// Todo = []task{{"get milk", false}, {"make breakfast", false}, {"wash car", false}, {"mow lawn", true}}
	Todo = []task{{"get milk", true}, {"make breakfast", true}, {"wash car", false}, {"mow lawn", true}, {"walk dog", false}, {"call mom", false}}
	// tpl, _ = tpl.ParseFiles("templates/*.html")
	// tpl, _ = tpl.ParseFiles("templates/index.html", "templates/groceries.html")

	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/todo", todoHandler)
	http.ListenAndServe(":80", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "groceries.html", g)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "todolist.html", Todo)
}
