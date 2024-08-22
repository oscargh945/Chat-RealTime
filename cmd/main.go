package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/oscargh945/go-Chat/domain/service"
	"github.com/oscargh945/go-Chat/infrastructure/http/Router"
	"github.com/oscargh945/go-Chat/infrastructure/http/handler"
	"github.com/oscargh945/go-Chat/infrastructure/postgresConfig"
	"github.com/oscargh945/go-Chat/infrastructure/repositories"
	"log"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar las .env")
	}
	postgres := postgresConfig.NewPostgres(ctx)
	postgres.InitPostgresDB()

	userRepository := repositories.NewUserRepository(postgres.Pool, ctx)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(*userService)

	Router.RouterInit(userHandler)
	Router.Init("localhost:8080")
}
