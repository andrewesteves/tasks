package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrewesteves/tasks/internal/handlers"
	"github.com/andrewesteves/tasks/internal/repositories"
)

func main() {
	mux := http.NewServeMux()

	taskRepository := repositories.NewTaskInMemory()
	taskHandler := handlers.NewTasks(taskRepository)

	mux.HandleFunc("/tasks", taskHandler.Actions)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tasks"))
	})

	fmt.Println("app is running")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Fatal(err)
	}
}
