package main

import (
	"alex/fishorder-api-v3/app2/server"
	"alex/fishorder-api-v3/app2/storage"
	"context"
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	debug := flag.Bool("d", false, "enables debug mode")
	flag.Parse()

	serverBindAddr := os.Getenv("BIND_ADDR")
	if serverBindAddr == "" {
		serverBindAddr = ":8080"
	}

	storageAddr := os.Getenv("STORAGE_ADDR")
	if storageAddr == "" {
		storageAddr = "localhost"
	}

	if *debug {
		log.SetLevel(logrus.DebugLevel)
	}
	log.Debug("Logger initialized")

	storageConfig := storage.Config{
		Host:     storageAddr,
		Port:     5432,
		User:     "postgres",
		Password: "password",
		DBname:   "postgres",
	}

	st, err := storage.New(context.Background(), &storageConfig)
	if err != nil {
		log.Fatalf("Failed to create a storage: %v", err)
	}
	log.Debug("Storage initialized")

	serverConfig := server.Config{
		BindAddr: serverBindAddr,
		Storage:  st,
	}

	serv, err := server.New(&serverConfig)
	if err != nil {
		log.Fatalf("Failed to create a server: %v", err)
	}
	log.Debug("Server initialized")

	if err := serv.Start(); err != nil {
		log.Fatalf("Failed to start a server: %v", err)
	}
}
