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

func NewMovie(log *log.Logger, movie_dao daos.MovieDao) *Movie {
	return &Movie{
		l:        log,
		moviedao: movie_dao,
	}
}

func (m *Movie) GetMovies(pagination *dtos.PaginationSpecifics) []dtos.Movie {
	var movies []dtos.Movie
	movies, err := m.moviedao.GetMoviesForHomePage(pagination)
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies
}

func (m *Movie) GetSuggestionMovies() []dtos.Movie {
	var movies []dtos.Movie
	movies, err := m.moviedao.GetMoviesForSuggestionPage()
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies
}

func (m *Movie) GetTrendingMovies() []dtos.Movie {
	var movies []dtos.Movie
	movies, err := m.moviedao.GetMoviesForTrendingPage()
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies
}
