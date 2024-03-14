package main

import (
	"log"
	"solidgate_test/internal/services"
	"solidgate_test/internal/transport"
	"solidgate_test/internal/transport/handlers"
)

func main() {
	service := services.NewService()
	handler := handlers.NewHandler(service)
	server := transport.NewRouter(handler)

	log.Printf("server listens on %v\n", server.Addr)
	log.Fatalln(server.ListenAndServe())
}
