package services

import (
	"fmt"
	"net/url"
	"strings"

	"movie-suggestions-api/config"
	"movie-suggestions-api/dtos"
	"movie-suggestions-api/utils/log"

	"github.com/PuerkitoBio/goquery"
)

type MovieRating struct {
	l *log.Logger
}

func NewMovieRating(l *log.Logger) *MovieRating {
	return &MovieRating{
		l: l,
	}
}

func (m *MovieRating) GetRating(name string) []dtos.Movie {
	name = url.QueryEscape(name)
	doc, err := goquery.NewDocument(fmt.Sprintf(config.IMDB_URL+config.SEARCH_URL, name))
	m.l.Info("Searching for movie: ", name)
	if err != nil {
		m.l.Fatal("Unable to search for movie in IMDB. Err: ", err)
	}

	var movies []dtos.Movie
	doc.Find(".findSection").Each(func(i int, s *goquery.Selection) {
		header := s.Find(".findSectionHeader").Text()
		if header == "Titles" {
			s.Find("tr").Each(func(i int, s *goquery.Selection) {
				movieName := s.Find("td").Text()
				m.l.Info("Movie name: ", movieName)
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					if i == 1 {
						href, _ := s.Find("a").Attr("href")
						movieName = strings.TrimSpace(movieName)
						doc, err := goquery.NewDocument(config.IMDB_URL + href)
						if err != nil {
							m.l.Fatal("Unable to scrape movie page in IMDB. Err: ", err)
						}
						m.l.Info("Scraping movie document: ", doc)
						rating, _ := doc.Find(".ratingValue").Find("strong").Attr("title")
						movies = append(movies, dtos.Movie{
							Name:   movieName,
							Rating: rating,
						})

						return
					} else {
						m.l.Info("Unable to find movie rating")
						return
					}
				})
			})
		}
	})
	m.l.Debug("Movies: ", movies)
	return movies
}
