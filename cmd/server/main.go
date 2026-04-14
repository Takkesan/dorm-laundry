package main

import (
	"log"
	"net/http"
	"os"

	"github.com/takke/dorm-laundry/internal/app"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("server init failed: %v", err)
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Handler()); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
