package main

import (
	beteratest "betera-test"
	"betera-test/pkg/handler"
	"betera-test/pkg/repository"
	"betera-test/pkg/service"
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(beteratest.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

	fmt.Println("Hello")
}
