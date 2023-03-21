package handlers

import (
	"net/http"

	"movie-suggestions-api/services"

	"github.com/julienschmidt/httprouter"
)

func setMovieRoutes(router *httprouter.Router) {
	router.GET("/movies/rating", MovieRating)
}

func MovieRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	m := services.NewMovieRating(rd.l)

	movies := m.GetRating(r.FormValue("name"))

	writeJSONStruct(movies, http.StatusOK, rd)
}
