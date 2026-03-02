package config

import (
	"flag"
	"log"
)

// type StorageType string

// const (
// 	StorageMongo StorageType = "mongo"
// 	StorageFiles StorageType = "files"
// )

type Config struct {
	TgBotToken  string
	StorageType string // "mongo" | "files"
	// StorageType           StorageType // "mongo" | "files"
	MongoConnectionString string
	StoragePath       string
}

func MustLoad() Config {
	tgBotTokenToken := flag.String("tg-bot-token", "", "token for access to telegram bot")
	storageType := flag.String("storage", "files", "storage type: mongo|files")
	mongoConnectionString := flag.String("mongo-connection-string", "", "connection string for MongoDB")
	storagePath := flag.String("storage-path", "", "base dir for file storage")

	flag.Parse()

	if *tgBotTokenToken == "" {
		log.Fatal("token is not specified")
	}

	switch *storageType {
	case "mongo":
		if *mongoConnectionString == "" {
			log.Fatal("mongo connection string is not specified")
		}
	case "files":
		if *storagePath == "" {
			log.Fatal("storage path is not specified")
		}
	default:
		log.Fatal("unknown storage type: " + *storageType)
	}

	return Config{
		TgBotToken:            *tgBotTokenToken,
		StorageType:           *storageType,
		MongoConnectionString: *mongoConnectionString,
		StoragePath:           *storagePath,
	}
}
