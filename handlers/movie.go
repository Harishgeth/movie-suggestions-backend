package handlers

import (
	"encoding/json"
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
}

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
	rd.l.Info("The content of scrollAnalyticsData", scrollAnalyticsData)
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
