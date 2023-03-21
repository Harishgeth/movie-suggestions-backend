package daos

import (
	"log"
	"movie-suggestions-api/config"
	"time"

	"gopkg.in/mgo.v2"
)

// get() returns a DB session
func get() *mgo.Session {
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout(config.ATLAS_URI, maxWait)
	log.Println("Here at get func")
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	log.Println("Outside session")

	return session
}
