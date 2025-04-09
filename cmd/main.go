package main

import (
	"log"
	"net/http"
	"test/handler"
	"test/repository"
	"test/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	handler := handler.NewUserHandler(uc)

	handler.RegisterRoutes(app)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
