package main

import (
	"log"

	"github.com/kvs4/static-server/Internal/config"
	"github.com/kvs4/static-server/Internal/server"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatalf("Didn`t read .env config: %v", err)
	}
	server := server.New(cfg)
	log.Printf("Starting server on port %s\n", cfg.Port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
