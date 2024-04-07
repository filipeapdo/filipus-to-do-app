package handler

import (
	"html/template"
	"net/http"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks = []Task{
	{1, "bla", false},
	{2, "ble", true},
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../tasks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, tasks)
}
