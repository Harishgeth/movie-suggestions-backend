package daos

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"movie-suggestions-api/config"
	dtos "movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(l *log.Logger) {

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(config.ATLAS_URI))
	if err != nil {
		l.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		l.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a new database and collection
	db := client.Database("movies")
	colName := "movies"
	colExists, err := collectionExists(ctx, db, colName)
	if err != nil {
		l.Fatal(err)
	}
	if !colExists {
		err = db.CreateCollection(ctx, colName)
		if err != nil {
			l.Fatal(err)
		}
	}

	// Insert data into the collection if it is empty
	count, err := db.Collection(colName).EstimatedDocumentCount(ctx)
	if err != nil {
		l.Fatal(err)
	}
	if count == 0 {
		// Read data from movies.json
		data, err := ioutil.ReadFile("./bootstrapped/movie.json")
		if err != nil {
			l.Fatal(err)
		}

		// Unmarshal data into a slice of movies
		var movies []dtos.Movie
		err = json.Unmarshal(data, &movies)
		if err != nil {
			l.Fatal(err)
		}
		l.Debug("Movies: ", movies)

		// Insert movies into the collection
		var documents []interface{}
		for _, movie := range movies {
			documents = append(documents, movie)
		}
		_, err = db.Collection(colName).InsertMany(ctx, documents)
		if err != nil {
			l.Fatal(err)
		}
	}
}

// collectionExists checks if a collection exists in a database
func collectionExists(ctx context.Context, db *mongo.Database, colName string) (bool, error) {
	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	for _, c := range collections {
		if c == colName {
			return true, nil
		}
	}
	return false, nil
}
