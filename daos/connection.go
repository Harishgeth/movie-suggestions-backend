package daos

import (
	"context"
	"movie-suggestions-api/config"
	"movie-suggestions-api/utils/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// // get() returns a DB session
// func get() *mgo.Session {
// 	maxWait := time.Duration(5 * time.Second)
// 	session, err := mgo.DialWithTimeout(config.ATLAS_URI, maxWait)
// 	log.Println("Here at get func")
// 	if err != nil {
// 		log.Fatalln(err)
// 		return nil
// 	}
// 	log.Println("Outside session")

// 	return session
// }

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
