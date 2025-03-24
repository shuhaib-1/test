package main

import (
	"fmt"
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

	fmt.Println("Server running on port: 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
