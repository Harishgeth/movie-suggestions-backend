package main

import (
	"fmt"
	"net/http"

	"movie-suggestions-api/config"
	"movie-suggestions-api/handlers"
	"movie-suggestions-api/utils/log"
)

func main() {

	l := log.NewLogger("")
	// connectToMongodb(l)
	fmt.Println(config.ATLAS_URI)

	l.Info("Port: ", config.PORT)
	http.ListenAndServe(":"+config.PORT, handlers.GetRouter())
}
