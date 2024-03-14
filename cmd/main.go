package main

import (
	"github.com/anton-uvarenko/solidgate_test/internal/services"
	"github.com/anton-uvarenko/solidgate_test/internal/transport"
	"github.com/anton-uvarenko/solidgate_test/internal/transport/handlers"
	"log"
)

func main() {
	service := services.NewService()
	handler := handlers.NewHandler(service)
	server := transport.NewRouter(handler)

	log.Printf("server listens on %v\n", server.Addr)
	log.Fatalln(server.ListenAndServe())
}
