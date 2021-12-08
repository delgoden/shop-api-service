package main

import (
	"log"

	"github.com/delgoden/shop-api-service/internal/api/handlers"
	"github.com/delgoden/shop-api-service/internal/api/server"
	"github.com/delgoden/shop-api-service/internal/service"
	"github.com/delgoden/shop-api-service/pkg/rest"
)

func main() {
	cli := &rest.BaseClient{}
	service := service.NewClient(cli)
	h := handlers.NewHandler(service)

	srv := new(server.Server)

	if err := srv.Run(":8001", h.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
