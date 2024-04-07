package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	handler "github.com/filipeapdo/filipus-to-do-app/api"
)

func main() {
	log.Println("FILIPUS To Do App")

	// TODO: remove at some point: only needed for dev purpose
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api/get-tasks", handler.GetTasks)

	http.HandleFunc("/api/add-task", handler.AddTask)

	http.HandleFunc("/api/del-task", handler.DelTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Starting server on %v\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
