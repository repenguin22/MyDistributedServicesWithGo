package main

import (
	"github.com/repenguin22/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
