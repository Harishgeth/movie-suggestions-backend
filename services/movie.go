package services

import (
	"encoding/json"
	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	elasticDao "movie-suggestions-api/elasticdao"
	"movie-suggestions-api/utils/log"
	"regexp"
	"strings"
)

type Movie struct {
	l          *log.Logger
	moviedao   daos.MovieDao
	elasticDao elasticDao.MovieElasticDao
}

func NewMovie(log *log.Logger, movie_dao daos.MovieDao, elasticDao elasticDao.MovieElasticDao) *Movie {
	return &Movie{
		l:          log,
		moviedao:   movie_dao,
		elasticDao: elasticDao,
	}
}

func (m *Movie) extractJsonFromLogString(logStr string) (*dtos.ScrollDataCaptured, error) {

	re := regexp.MustCompile(`\{.*\}`)
	jsonStr := re.FindString(logStr)

	var jsonData dtos.ScrollDataCaptured
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		m.l.Error("Error parsing JSON:", err)
		return nil, err
	}

	scrollBytes, _ := json.Marshal(jsonData)
	m.l.Info("Parsed scrolldata content:", string(scrollBytes))
	return &jsonData, nil
}

func (m *Movie) GetMovies(pagination *dtos.PaginationSpecifics) []dtos.Movie {
	var movies []dtos.Movie
	movies, err := m.moviedao.GetMoviesForHomePage(pagination)
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies
}

func (m *Movie) GetSuggestionMovies(pagination *dtos.PaginationSpecifics, userid string) ([]dtos.Movie, error) {

	suggestedPostIDs, err := m.elasticDao.GetSuggestionsPageSortOrder(userid, pagination.Page, pagination.Limit)
	if err != nil {
		m.l.Error("Failed to get the movies sort order from Elasticsearch:", err)
		return nil, err
	}
	movies, err := m.moviedao.GetMoviesByPostIDs(suggestedPostIDs)
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
		return nil, err
	}
	return movies, nil
}

func (m *Movie) GetTrendingMovies(pagination *dtos.PaginationSpecifics) ([]dtos.Movie, error) {

	suggestedPostIDs, err := m.elasticDao.GetTrendingPageSortOrder(pagination.Page, pagination.Limit)
	if err != nil {
		m.l.Error("Failed to get the movies sort order from Elasticsearch:", err)
		return nil, err
	}
	movies, err := m.moviedao.GetMoviesByPostIDs(suggestedPostIDs)
	if err != nil {
		m.l.Error("Faced an error while getting from db", err)
	}
	return movies, nil
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
				moviedatapoint, err := m.extractJsonFromLogString(dockerLogs[i].Message)
				if err != nil {
					m.l.Error("Error while extracting json from log string", err)
					return err
				}
				err = m.elasticDao.PutDataInElasticSearch(moviedatapoint)
				if err != nil {
					m.l.Error("Error while putting data in elastic search", err)
					return err
				}
				m.l.Info("Successfully put data in elastic search")
			}
		}
	}
	return nil
}
