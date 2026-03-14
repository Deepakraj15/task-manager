package main

import (
	"fmt"
	"net/http"

	"github.com/Deepakraj15/task-manager/api/handlers"
	"github.com/Deepakraj15/task-manager/internal/algo"
	"github.com/Deepakraj15/task-manager/internal/constants"
	"github.com/Deepakraj15/task-manager/internal/logger"
	"github.com/go-chi/chi/v5"
)

/*
*

	Main function chosen for execution entry point

*
*/
func main() {
	log, logError := logger.GetLogger(constants.APPLICATION_LOG)
	if logError != nil {
		fmt.Println("Error occurred while initializing logs ERROR -> " + logError.Error())
	}
	log.Info("Starting Task manager Application....")
	router := chi.NewRouter()
	handlers.TaskHandlers(router)
	log.Info("Application is ready to accept traffic...")
	go algo.StartWheel()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Error("Error occurred while starting application", err)
	}
}
