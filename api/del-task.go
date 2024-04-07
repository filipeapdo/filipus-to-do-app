package handler

import (
	"net/http"
	"slices"
	"strconv"
)

func remove(slice []Task, s int) []Task {
	idx := slices.IndexFunc(slice, func(t Task) bool { return t.ID == s })
	return append(slice[:idx], slice[idx+1:]...)
}

func DelTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks = remove(tasks, id)
	http.Redirect(w, r, "/api/get-tasks", http.StatusFound)
}
