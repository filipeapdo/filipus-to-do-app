package main

import (
	"html/template"
	"log"
	"net/http"
	"slices"
	"strconv"

	handler "github.com/filipeapdo/filipus-to-do-app/api"
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

func remove(slice []Task, s int) []Task {
	idx := slices.IndexFunc(slice, func(t Task) bool { return t.ID == s })
	log.Println(idx)
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {
	log.Println("FILIPUS To Do App")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("tasks.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, tasks)
	})

	http.HandleFunc("/api/tasks", handler.Tasks)

	http.HandleFunc("/add-task", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			title := r.FormValue("title")
			tasks = append(tasks, Task{len(tasks) + 1, title, false})
			log.Println(tasks)
			http.Redirect(w, r, "/tasks", http.StatusFound)
		}
	})

	http.HandleFunc("/delete-task", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tasks = remove(tasks, id)
		log.Println(tasks)
		http.Redirect(w, r, "/tasks", http.StatusFound)
	})

	log.Println("Starting server on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
