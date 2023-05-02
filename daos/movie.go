package daos

import (
	dtos "movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieDao interface {
	GetMoviesForHomePage(pagination *dtos.PaginationSpecifics) ([]dtos.Movie, error)
	GetMoviesByPostIDs([]string) ([]dtos.Movie, error)
}
type Movie struct {
	l *log.Logger
}

func GetMovieDao(l *log.Logger) MovieDao {
	return Movie{
		l: l,
	}
}

const DBNAME = "movies"

const DOCNAME = "movies"

func (dao Movie) GetMoviesForHomePage(pagination *dtos.PaginationSpecifics) ([]dtos.Movie, error) {
	client, ctx := get(dao.l)
	defer client.Disconnect(ctx)

	res := []dtos.Movie{}
	coll := client.Database("movies").Collection("movies")
	if coll == nil {
		dao.l.Error("Collection is nil")
		panic("Collection is nil")
	}
	findOptions := options.Find()

	findOptions.SetSkip(int64(pagination.Skip))
	findOptions.SetLimit(int64(pagination.Limit))

	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}

	return res, nil

}

func (dao Movie) GetMoviesByPostIDs(postIDs []string) ([]dtos.Movie, error) {
	client, ctx := get(dao.l)
	defer client.Disconnect(ctx)

	res := []dtos.Movie{}
	coll := client.Database("movies").Collection("movies")
	if coll == nil {
		dao.l.Error("Collection is nil")
		panic("Collection is nil")
	}

	filter := bson.M{"movie_id": bson.M{"$in": postIDs}}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}

	resMap := make(map[string]dtos.Movie)
	for i := 0; i < len(res); i++ {
		resMap[res[i].MovieID] = res[i]
	}
	sortedRes := []dtos.Movie{}
	for i := 0; i < len(postIDs); i++ {
		if movie, ok := resMap[postIDs[i]]; ok {
			sortedRes = append(sortedRes, movie)
		}
	}

	return sortedRes, nil
}
