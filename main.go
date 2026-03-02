package main

import (
	"log"
	"time"

	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/config"
	"read-adviser-bot/storage"

	eventConsumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
	"read-adviser-bot/storage/mongo"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	cfg := config.MustLoad()

	var storage storage.Storage

	switch cfg.StorageType {
	case "mongo":
		storage = mongo.New(cfg.MongoConnectionString, 10*time.Second)
	case "files":
		storage = files.New(storagePath)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
