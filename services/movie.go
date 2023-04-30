package services

import (
	"encoding/json"
	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"
	"regexp"
	"strings"
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

func (m *Movie) extractJsonFromLogString(logStr string) *dtos.ScrollDataCaptured {

	re := regexp.MustCompile(`\{.*\}`)
	jsonStr := re.FindString(logStr)

	var jsonData dtos.ScrollDataCaptured
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		m.l.Error("Error parsing JSON:", err)
		return nil
	}

	scrollBytes, _ := json.Marshal(jsonData)
	m.l.Info("Parsed scrolldata content:", string(scrollBytes))
	return &jsonData
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

func (m *Movie) FilterAndDigestLogIntoElasticSearch(log_entity string) error {
	if strings.Contains(log_entity, "scrollAnalyticsData:") {
		var dockerLogs []dtos.DockerLog
		err := json.Unmarshal([]byte(log_entity), &dockerLogs)
		if err != nil {
			m.l.Error("Error while decoding the scroll data", err)
			return err
		}
		for i := 0; i < len(dockerLogs); i++ {
			if strings.Contains(dockerLogs[i].Message, "scrollAnalyticsData:") {
				_ = m.extractJsonFromLogString(dockerLogs[i].Message)
			}
		}
	}
	return nil
}
