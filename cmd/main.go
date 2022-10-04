package main

import (
	beteratest "betera-test"
	"betera-test/pkg/client"
	"betera-test/pkg/handler"
	"betera-test/pkg/repository"
	"betera-test/pkg/service"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	client := client.NewClient(os.Getenv("APOD_URL"))
	service := service.NewService(repos, client)
	handlers := handler.NewHandler(service)

	srv := new(beteratest.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

	fmt.Println("Hello")
}
