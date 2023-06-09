package main

import (
	"net/http"

	"movie-suggestions-api/config"
	"movie-suggestions-api/daos"
	elasticDao "movie-suggestions-api/elasticdao"
	"movie-suggestions-api/handlers"
	"movie-suggestions-api/utils/log"
)

func main() {

	l := log.NewLogger("")
	daos.Init(l)
	elasticDao := elasticDao.GetMovieDao(l)
	err := elasticDao.CreateMovieIndexIfNotExists()
	if err != nil {
		l.Fatal("Error creating movie index: ", err)
	}
	// connectToMongodb(l)
	l.Info("Port: ", config.PORT)
	http.ListenAndServe(":"+config.PORT, handlers.GetRouter())
}
