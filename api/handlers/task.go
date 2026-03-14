package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Deepakraj15/task-manager/internal/algo"
	"github.com/Deepakraj15/task-manager/internal/models"
	"github.com/go-chi/chi/v5"
)

func TaskHandlers(router chi.Router) {

	router.Route("/tasks", func(router chi.Router) {

		// to get all tasks
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello it works")
		})

		router.Post("/submitTask", func(w http.ResponseWriter, r *http.Request) {
			var task models.Task

			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fmt.Println(task.TName)
			fmt.Println(task.TDescription)
			fmt.Println(task.RunAt)

			algo.SubmitTask(task)

			w.WriteHeader(http.StatusOK)
		})

		router.Delete("/deleteTask", func(w http.ResponseWriter, r *http.Request) {
			chi.URLParam(r, "taskId")
		})
	})
}
