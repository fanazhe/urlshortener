package main

import (
	"log"
	"urlshortener/config"
	"urlshortener/handler"
	"urlshortener/router"
	"urlshortener/storage"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.New(config.Redis.Host, config.Redis.Port, config.Redis.DB)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler(db)
	r := router.New(h, config)

	r.Logger.Fatal(r.Start(config.Server.Host + ":" + config.Server.Port))
}
