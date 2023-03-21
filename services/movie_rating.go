package services

import (
	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"
)

type MovieRating struct {
	l        *log.Logger
	moviedao daos.MovieDao
}

func NewMovieRating(l *log.Logger, moviedao daos.MovieDao) *MovieRating {
	return &MovieRating{
		l:        l,
		moviedao: moviedao,
	}
}

func (m *MovieRating) GetRating(name string) []dtos.Movie {
	var movies []dtos.Movie

	return movies
}
