package main

import (
	"fmt"
	"net/http"

	"movie-suggestions-api/config"
	"movie-suggestions-api/handlers"
	"movie-suggestions-api/utils/log"
)

// func connectToMongodb(l *log.Logger) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(config.ATLAS_URI))
// 	if err != nil {
// 		l.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		l.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	/*
// 		List databases
// 	*/
// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		l.Fatal(err)
// 	}
// 	fmt.Println(databases)
// }

func main() {

	l := log.NewLogger("")
	// connectToMongodb(l)
	fmt.Println(config.ATLAS_URI)

	l.Info("Port: ", config.PORT)
	http.ListenAndServe(":"+config.PORT, handlers.GetRouter())
}
