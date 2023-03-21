package daos

import (
	dtos "movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"
)

type MovieDao interface {
	GetMoviesForHomePage() ([]dtos.Movie, error)
}
type Movie struct {
}

func GetMovieDao(l *log.Logger) MovieDao {
	return Movie{}
}

const DBNAME = "movies"

const DOCNAME = "movies"

func (dao Movie) GetMoviesForHomePage() ([]dtos.Movie, error) {
	db := get()
	defer db.Close()

	res := []dtos.Movie{}

	if err := db.DB(DBNAME).C(DOCNAME).Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil

}
