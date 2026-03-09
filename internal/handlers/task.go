package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TaskHandlers(router chi.Router) {

	router.Route("/tasks", func(router chi.Router) {

		// to get all tasks
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello it works")
		})

		router.Post("/submitTask", func(w http.ResponseWriter, r *http.Request) {})

		router.Delete("/deleteTask", func(w http.ResponseWriter, r *http.Request) {
			chi.URLParam(r, "taskId")
		})
	})
}
