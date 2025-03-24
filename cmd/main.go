package main

import (
	"log"
	"net/http"
	"test/handler"
	"test/repository"
	"test/usecase"
)

func main() {

	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	handler := handler.NewUserHandler(uc)

	handler.RegisterRoutes()

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
