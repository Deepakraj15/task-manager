package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deepakraj15/task-manager/internal/algo"
	"github.com/Deepakraj15/task-manager/internal/models"
	"github.com/go-chi/chi/v5"
)

func TaskHandlers(router chi.Router) {

	router.Route("/tasks", func(router chi.Router) {

		// to get all tasks
		router.Get("/getPendingTasks", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			tasks := models.GetSecondsWheel().GetAllPendingTasks()

			err := json.NewEncoder(w).Encode(tasks)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		})
		router.Post("/submitTask", func(w http.ResponseWriter, r *http.Request) {
			var task models.Task

			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			algo.SubmitTask(task)

			w.WriteHeader(http.StatusOK)
		})

	})
}
