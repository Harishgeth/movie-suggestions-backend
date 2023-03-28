package daos

import (
	dtos "movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"

	"go.mongodb.org/mongo-driver/bson"
)

type MovieDao interface {
	GetMoviesForHomePage() ([]dtos.Movie, error)
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

func (dao Movie) GetMoviesForHomePage() ([]dtos.Movie, error) {
	client, ctx := get(dao.l)
	defer client.Disconnect(ctx)

	res := []dtos.Movie{}
	coll := client.Database("movies").Collection("movies")
	if coll == nil {
		dao.l.Error("Collection is nil")
		panic("Collection is nil")
	}

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &res); err != nil {
		panic(err)
	}

	// if err := db.DB(DBNAME).C(DOCNAME).Find(nil).All(&res); err != nil {
	// 	return nil, err
	// }

	return res, nil

}
