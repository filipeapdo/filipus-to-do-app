package handler

import (
	"html/template"
	"log"
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

const tmplTasks = `
{{range .}}
<div>
  <span>{{.ID}} - {{.Title}}</span>
  <form
    hx-post="/delete-task"
    hx-target="#todo-list"
    hx-confirm="Deleting a task?"
  >
    <input type="hidden" name="id" value="{{.ID}}" />
    <button type="submit">Delete</button>
  </form>
</div>
{{end}}
`

func Tasks(w http.ResponseWriter, r *http.Request) {
	log.Println("from /api/taks...")
	tmpl, err := template.New("tasks").Parse(tmplTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, tasks)
}
