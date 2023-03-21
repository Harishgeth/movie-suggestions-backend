package handlers

import (
	"net/http"

	"runtime/debug"

	"movie-suggestions-api/config"
	"movie-suggestions-api/config/globals"
	"movie-suggestions-api/utils/log"

	"github.com/julienschmidt/httprouter"
)

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	globals.Logger = log.NewLogger("")

	router := httprouter.New()
	router.PanicHandler = PanicHandler
	setMovieRoutes(router)

	return router
}

func tokenAuthentication(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userName, _, ok := r.BasicAuth()
		if (ok && userName == config.AUTH_TOKEN) || r.FormValue("token") == config.AUTH_TOKEN {
			h(w, r, ps)
			return
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	globals.Logger.Fatal("Recovering from panic, Reason: ", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
