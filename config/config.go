package config

import "os"

var (
	AUTH_TOKEN     = os.Getenv("AUTH_TOKEN")
	PORT           = os.Getenv("PORT")
	ATLAS_URI      = os.Getenv("ATLAS_URI")
	ELASTIC_URI    = os.Getenv("ELASTIC_URI")
	MOVIE_INDEX    = "movie_index"
	ES_TIME_FORMAT = "2006-01-02 15:04:05"
)
