package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>FILIPUS To Do App</h1><p>(from index.go)</p>")
}
