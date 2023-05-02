package daos

import (
	"context"
	"movie-suggestions-api/config"
	"movie-suggestions-api/utils/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func get(l *log.Logger) (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.ATLAS_URI))
	if err != nil {
		l.Fatal(err)
		// return nil, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		l.Fatal(err)
	}
	return client, ctx
}
