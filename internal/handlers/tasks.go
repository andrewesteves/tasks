package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/andrewesteves/tasks/internal/entities"
	"github.com/andrewesteves/tasks/internal/repositories"
	"github.com/google/uuid"
)

type Tasks struct {
	taskRepository repositories.Task
}

func NewTasks(taskRepository repositories.Task) *Tasks {
	return &Tasks{
		taskRepository: taskRepository,
	}
}

func (t *Tasks) Actions(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	switch action {
	case "create":
		t.create(w, r)
	case "complete":
		t.complete(w, r)
	default:
		t.index(w, r)
	}
}

func (t *Tasks) index(w http.ResponseWriter, r *http.Request) {
	views := []string{
		"./internal/views/layout.tpl.html",
		"./internal/views/tasks/index.tpl.html",
	}

	ts, err := template.ParseFiles(views...)
	if err != nil {
		log.Println(err)
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	tasks := t.taskRepository.Get()
	if err := ts.ExecuteTemplate(w, "layout", tasks); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
	}
}

func (t *Tasks) create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			return
		}

		t.taskRepository.Put(
			entities.Task{
				ID:    uuid.NewString(),
				Title: r.FormValue("title"),
			},
		)

		http.Redirect(w, r, "/tasks", http.StatusFound)
	}
}

func (t *Tasks) complete(w http.ResponseWriter, r *http.Request) {
	t.taskRepository.Delete(r.URL.Query().Get("id"))
	http.Redirect(w, r, "/tasks", http.StatusFound)
}
