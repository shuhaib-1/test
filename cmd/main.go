package main

import (
	"net/http"
	"test/handler"
	"test/logger"
	"test/repository"
	"test/usecase"
)

func main() {
	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	handler := handler.NewUserHandler(uc)

	// Register routes first
	handler.RegisterRoutes()

	// Create a new ServeMux and wrap it with the logger middleware
	mux := http.DefaultServeMux
	wrappedMux := logger.HTTPLogger(mux)

	logger.Info("The project has started at :3000")
	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		logger.Error("Error starting server: ", err)
	}
}
