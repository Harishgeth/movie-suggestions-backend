package services

import (
	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"
)

type Movie struct {
	l        *log.Logger
	moviedao daos.MovieDao
}

func NewMovie(l *log.Logger, moviedao daos.MovieDao) *Movie {
	return &Movie{
		l:        l,
		moviedao: moviedao,
	}
}

func (m *Movie) GetMovies() []dtos.Movie {
	var movies []dtos.Movie
	movies, err := m.moviedao.GetMoviesForHomePage()
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies
}
