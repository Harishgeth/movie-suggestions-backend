package config

import "os"

var (
	AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
	PORT       = os.Getenv("PORT")
	IMDB_URL   = "http://www.imdb.com"
	SEARCH_URL = "/find?ref_=nv_sr_fn&q=%s&s=all"
)
