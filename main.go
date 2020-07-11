package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/kevinburke/twilio-go"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	var storage ConversationStorageAdapter
	if config.Storage.Adapter == RedisStorageAdapterType {
		storage = NewRedisStorageAdapter(redis.NewClient(
			&redis.Options{
				Addr:     config.Storage.Host,
				Password: config.Storage.Password,
				DB:       config.Storage.DB,
			},
		))
	} else if config.Storage.Adapter == MemoryStorageAdapterType {
		storage = NewMemoryStorageAdapter()
	} else {
		panic("Storage adapter in config was not valid")
	}

	twilioClient := twilio.NewClient(config.Twilio.SID, config.Twilio.Token, nil)

	app := &TextScreen{
		Storage: storage,
		Config:  config,
		Twilio:  twilioClient,
	}

	http.Handle("/twilio", app)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), nil)
}
