package handlers

import (
	"net/http"

	"movie-suggestions-api/daos"
	"movie-suggestions-api/services"

	"github.com/julienschmidt/httprouter"
)

func setMovieRoutes(router *httprouter.Router) {
	router.GET("/home-page", GetMovies)
}

// func MovieRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	rd := logAndGetContext(w, r)
// 	m := services.NewMovieRating(rd.l)

// 	movies := m.GetRating(r.FormValue("name"))

// 	writeJSONStruct(movies, http.StatusOK, rd)
// }

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	rd.l.Info("Here right now")
	m := services.NewMovie(rd.l, daos.GetMovieDao(rd.l))
	movies := m.GetMovies()
	writeJSONStruct(movies, http.StatusOK, rd)
}
