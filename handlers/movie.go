package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"movie-suggestions-api/daos"
	"movie-suggestions-api/dtos"
	elasticDao "movie-suggestions-api/elasticdao"
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

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		rd.l.Info("The page number is not provided")
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)

	LimitStr := r.URL.Query().Get("limit")
	if LimitStr == "" {
		rd.l.Info("The count is not provided")
		LimitStr = "2"
	}
	limit, _ := strconv.Atoi(LimitStr)

	skip := (page - 1) * limit

	pagination := &dtos.PaginationSpecifics{
		Page:  page,
		Limit: limit,
		Skip:  skip,
	}

	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l), elasticDao.GetMovieDao(rd.l))
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
		return
	}
	scrollBytes, _ := json.Marshal(scrollAnalyticsData)
	rd.l.Info("The content of scrollAnalyticsData:", string(scrollBytes))
	writeJSONMessage("Successfully recorded the scroll data", MSG, http.StatusOK, rd)

}

func GetSuggestionMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		rd.l.Info("The page number is not provided")
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)

	LimitStr := r.URL.Query().Get("limit")
	if LimitStr == "" {
		rd.l.Info("The limit is not provided")
		LimitStr = "2"
	}
	limit, _ := strconv.Atoi(LimitStr)

	skip := (page - 1) * limit

	pagination := &dtos.PaginationSpecifics{
		Page:  page,
		Limit: limit,
		Skip:  skip,
	}

	user := r.URL.Query().Get("userid")
	if user == "" {
		rd.l.Info("The User ID is not provided")
		writeJSONMessage("Please provide a valid userID", ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l), elasticDao.GetMovieDao(rd.l))
	movies, err := m.GetSuggestionMovies(pagination, user)
	if err != nil {
		writeJSONMessage("Something broke while getting the data", ERR_MSG, http.StatusInternalServerError, rd)
	}
	rd.l.Info("The content of movies", movies)
	writeJSONStruct(movies, http.StatusOK, rd)
}

func GetTrendingMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		rd.l.Info("The page number is not provided")
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)

	LimitStr := r.URL.Query().Get("limit")
	if LimitStr == "" {
		rd.l.Info("The limit is not provided")
		LimitStr = "2"
	}
	limit, _ := strconv.Atoi(LimitStr)

	skip := (page - 1) * limit

	pagination := &dtos.PaginationSpecifics{
		Page:  page,
		Limit: limit,
		Skip:  skip,
	}
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l), elasticDao.GetMovieDao(rd.l))
	movies, err := m.GetTrendingMovies(pagination)
	if err != nil {
		writeJSONMessage("Something broke while getting the data", ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	rd.l.Info("The content of movies", movies)
	writeJSONStruct(movies, http.StatusOK, rd)
}

func CaptureDataToIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONMessage("Failed to record the data", ERR_MSG, http.StatusBadRequest, rd)
	}
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l), elasticDao.GetMovieDao(rd.l))
	err = m.FilterAndDigestLogIntoElasticSearch(string(b))
	if err != nil {
		writeJSONMessage("Failed to record the data", ERR_MSG, http.StatusBadRequest, rd)
	}
	writeJSONMessage("Success", MSG, http.StatusOK, rd)

}
