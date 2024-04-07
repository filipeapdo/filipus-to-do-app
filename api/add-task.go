package handler

import "net/http"

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		tasks = append(tasks, Task{len(tasks) + 1, title, false})
		http.Redirect(w, r, "/api/get-tasks", http.StatusFound)
	}
}
