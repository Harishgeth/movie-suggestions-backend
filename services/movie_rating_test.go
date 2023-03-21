package services

import (
	"testing"

	"movie-suggestions-api/utils/log"
)

func TestGetNameRating1(t *testing.T) {
	l := log.NewLogger("")
	m := NewMovieRating(l)
	movies := m.GetRating("deadpool")

	if len(movies) == 0 {
		t.Error("Unable to get movie name and rating")
		t.FailNow()
	}
}

func TestGetNameRating2(t *testing.T) {
	l := log.NewLogger("")
	m := NewMovieRating(l)
	movies := m.GetRating("just like heaven")

	if len(movies) == 0 {
		t.Error("Unable to get movie name and rating")
		t.FailNow()
	}

}
