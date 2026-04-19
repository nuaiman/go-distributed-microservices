package main

import (
	"context"
	"log"
	"logger/data"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RPCServer struct {
	DB *mongo.Client
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, response *string) error {
	entry := data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	}

	collection := r.DB.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println(err)
		return err
	}

	*response = "Logged via RPC"
	return nil
}
