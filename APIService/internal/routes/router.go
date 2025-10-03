package routes

import (
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/handlers"
)

func SetupRouters(handler *handlers.Handlers) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /create", handler.CreateTask)
	router.HandleFunc("GET /list", handler.GetTask)
	router.HandleFunc("DELETE /delete", handler.DeleteTask)
	router.HandleFunc("PUT /done", handler.UpdateTask)

	return router

}
