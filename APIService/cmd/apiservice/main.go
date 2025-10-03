package main

import (
	"fmt"
	"net/http"

	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/handlers"
	"github.com/ZeroZeroZerooZeroo/ChecklistApp/apiservice/internal/routes"
)

func main() {
	// TODO: инициализация конфига

	// TODO: инициализация хендлера
	handler := handlers.NewHandler()

	router := routes.SetupRouters(handler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)

	// TODO: запуск приложения
}
