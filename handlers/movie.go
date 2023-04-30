package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	"movie-suggestions-api/services"

	"github.com/julienschmidt/httprouter"
)

func setMovieRoutes(router *httprouter.Router) {
	router.GET("/home-page", GetMovies)
	router.POST("/record-scroll", RecordScroll)
	router.GET("/suggestion-page", GetSuggestionMovies)
	router.GET("/trending-page", GetTrendingMovies)
	router.POST("/capture-data-to-index", CaptureDataToIndex)
}

// type Log struct {
// 	ContainerCreatedAt string `json:"container_created_at"`
// 	ContainerID        string `json:"container_id"`
// 	ContainerName      string `json:"container_name"`
// 	Host               string `json:"host"`
// 	Image              string `json:"image"`
// 	Label              struct {
// 		ConfigHash         string `json:"com.docker.compose.config-hash"`
// 		ContainerNumber    string `json:"com.docker.compose.container-number"`
// 		DependsOn          string `json:"com.docker.compose.depends_on"`
// 		ComposeImage       string `json:"com.docker.compose.image"`
// 		OneOff             string `json:"com.docker.compose.oneoff"`
// 		Project            string `json:"com.docker.compose.project"`
// 		ProjectConfigFiles string `json:"com.docker.compose.project.config_files"`
// 		ProjectWorkingDir  string `json:"com.docker.compose.project.working_dir"`
// 		Service            string `json:"com.docker.compose.service"`
// 		Version            string `json:"com.docker.compose.version"`
// 	} `json:"label"`
// 	Message    string `json:"message"`
// 	SourceType string `json:"source_type"`
// 	Stream     string `json:"stream"`
// 	Timestamp  string `json:"timestamp"`
// }

// func MovieRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	rd := logAndGetContext(w, r)
// 	m := services.NewMovieRating(rd.l)

// 	movies := m.GetRating(r.FormValue("name"))

// 	writeJSONStruct(movies, http.StatusOK, rd)
// }

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		rd.l.Info("The page number is not provided")
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)

	LimitStr := r.URL.Query().Get("limit")
	if pageStr == "" {
		rd.l.Info("The page number is not provided")
		LimitStr = "2"
	}
	limit, _ := strconv.Atoi(LimitStr)

	skip := (page - 1) * limit

	pagination := &dtos.PaginationSpecifics{
		Page:  page,
		Limit: limit,
		Skip:  skip,
	}

	rd.l.Info("Here right now")
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l))
	movies := m.GetMovies(pagination)
	rd.l.Info("The content of movies", movies)
	writeJSONStruct(movies, http.StatusOK, rd)
}

func RecordScroll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	rd.l.Info("Here right now")
	decoder := json.NewDecoder(r.Body)
	var scrollAnalyticsData dtos.ScrollDataCaptured
	err := decoder.Decode(&scrollAnalyticsData)
	if err != nil {
		rd.l.Error("Error while decoding the scroll data", err)
		writeJSONMessage("Failed to record the data", ERR_MSG, http.StatusBadRequest, rd)
	}
	scrollBytes, _ := json.Marshal(scrollAnalyticsData)
	rd.l.Info("The content of scrollAnalyticsData:", string(scrollBytes))
	writeJSONMessage("Successfully recorded the scroll data", MSG, http.StatusOK, rd)

}

func GetSuggestionMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	rd.l.Info("Here right now")
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l))
	movies := m.GetSuggestionMovies()
	rd.l.Info("The content of movies", movies)
	writeJSONStruct(movies, http.StatusOK, rd)

	// decoder := json.NewDecoder(r.Body)
	// var scrollAnalyticsData dtos.Movie
	// err := decoder.Decode(&scrollAnalyticsData)
	// if err != nil {
	// 	rd.l.Error("Error while projecting scroll data ", err)
	// 	writeJSONMessage("Failed to display the scroll data", ERR_MSG, http.StatusBadRequest, rd)
	// }
}

func GetTrendingMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	rd.l.Info("Here right now")
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l))
	movies := m.GetTrendingMovies()
	rd.l.Info("The content of movies", movies)
	writeJSONStruct(movies, http.StatusOK, rd)
}

func CaptureDataToIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONMessage("Failed to record the data", ERR_MSG, http.StatusBadRequest, rd)
	}
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l))
	err = m.FilterAndDigestLogIntoElasticSearch(string(b))
	if err != nil {
		writeJSONMessage("Failed to record the data", ERR_MSG, http.StatusBadRequest, rd)
	}
}
