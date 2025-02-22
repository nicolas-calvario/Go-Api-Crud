package main

import (
	"fmt"
	"github.com/nicolas-calvario/Go-Api-Crud/internal/handler"
	"github.com/nicolas-calvario/Go-Api-Crud/internal/repository"
	"github.com/nicolas-calvario/Go-Api-Crud/internal/router"
	"github.com/nicolas-calvario/Go-Api-Crud/internal/service"
	"github.com/nicolas-calvario/Go-Api-Crud/pkg/database"
	"net/http"
)

func main() {
	db := database.ConnectDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := router.Routes(userHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
