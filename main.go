package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rzldimam28/tabungan-api/internal/api"
	"github.com/rzldimam28/tabungan-api/internal/config"
	"github.com/rzldimam28/tabungan-api/internal/database"
	"github.com/rzldimam28/tabungan-api/internal/repository"
	"github.com/rzldimam28/tabungan-api/internal/service"
)

func main() {

	cnf := config.NewConfig()
	dbConnection := database.NewDatabaseConnection(cnf)

	accountRepository := repository.NewAccount()
	mutationRepository := repository.Newmutation()

	accountService := service.NewAccount(dbConnection, accountRepository, mutationRepository)
	mutationService := service.NewMutation(dbConnection, mutationRepository)

	app := fiber.New()
	api.NewAccount(app, accountService, mutationService)

	log.Fatal(app.Listen(cnf.Server.Host + ":" + cnf.Server.Port))

}